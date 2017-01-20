// Copyright 2017, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"bytes"
	"database/sql"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"path/filepath"
	"regexp"
	"strings"
	"syscall"
	"time"

	"runtime"

	"reflect"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/golang/glog"
	"github.com/youtube/vitess/go/exit"
	"github.com/youtube/vitess/go/vt/dbconfigs"
	"github.com/youtube/vitess/go/vt/logutil"
	"github.com/youtube/vitess/go/vt/proto/topodata"
	"github.com/youtube/vitess/go/vt/servenv"
	"github.com/youtube/vitess/go/vt/sqlparser"
	"github.com/youtube/vitess/go/vt/tabletmanager/tmclient"
	"github.com/youtube/vitess/go/vt/tabletserver"
	"github.com/youtube/vitess/go/vt/topo"
	"github.com/youtube/vitess/go/vt/topo/topoproto"
	"github.com/youtube/vitess/go/vt/topotools"
	"github.com/youtube/vitess/go/vt/vtgate/engine"
	"github.com/youtube/vitess/go/vt/vtgate/planbuilder"
	"github.com/youtube/vitess/go/vt/vtgate/vindexes"
	"github.com/youtube/vitess/go/vt/wrangler"
	"golang.org/x/net/context"
)

var (
	tailBytes      = flag.Int64("tailBytes", 10000, "bytes to parse of input file, starting at the end")
	waitTime       = flag.Duration("wait-time", 1*time.Minute, "time to wait while parsing lines")
	cell           = flag.String("parse-cell", "", "cell to execute against")
	keyspace       = flag.String("parse-keyspace", "", "keyspace to execute against")
	connUser       = flag.String("parse-conn-user", "root", "User to connect to db, i.e. root")
	connPassword   = flag.String("parse-conn-password", "", "Password to connect to db")
	connHost       = flag.String("parse-conn-host", "localhost:3306", "Host and port to connect to db, i.e. localhost:3306")
	connSocket     = flag.String("parse-conn-socket", "", "Unix socket file to connect on, i.e. /var/run/mysql.sock")
	connFile       = flag.String("parse-conn-file", "", "Path to file containing connect info for db. Should contain at least one each of user=,host=,password=, on separate lines")
	outFile        = flag.String("parse-out-file", "", "Path to a file where any unparseable sql queries and their corresponding error output will be sent")
	compareExplain = flag.Bool("parse-compare-explain", true, "Whether to run explains of queries against the db to check that the parsed queries make sense")
	verbosity      = flag.Int("parse-verbosity", 1, "How verbose should logging be? 1 == just final stats, 2 == ignored errors, 3 == all")
	alias          = flag.String("alias", "", "Tablet alias to use in initializing tablet before parsing")
	hostname       = flag.String("hostname", "", "Hostname to use in initializing tablet before parsing")
	justInitialize = flag.Bool("parse-just-init", false, "If set, will initialize the tablet and schema then exit")

	ignoredErrorPatterns = []*regexp.Regexp{
		regexp.MustCompile("(keyspace) (\\w+) not found in vschema"),
		regexp.MustCompile("(symbol) (@@[\\w.]+) not found"),
		regexp.MustCompile("(Table) ('\\w+\\.\\w+') doesn't exist"),
	}

	handledErrorPatterns = []handledError{
		{regexp.MustCompile("(?i)ON DUPLICATE KEY UPDATE.*[^=]+=\\s*VALUES\\([^)]+\\)"), "INSERT ... ON DUPLICATE KEY UPDATE ... = VALUES(...)"},
		{regexp.MustCompile("(?i)ON DUPLICATE KEY UPDATE.*[^=]+=\\s*CASE WHEN"), "INSERT ... ON DUPLICATE KEY UPDATE ... = CASE WHEN ..."},
		{regexp.MustCompile("(?i)SELECT.* UNION .*SELECT"), "SELECT.. UNION"},
		{regexp.MustCompile("(?i)DELETE .+ FROM"), "Multi-Delete"},
		{regexp.MustCompile("(?i)REPLACE\\(.*\\)"), "REPLACE() function"},
		{regexp.MustCompile("(?i)UNHEX\\(.*\\)"), "UNHEX() function"},
		{regexp.MustCompile("(?i)CURRENT_TIMESTAMP|utc_timestamp|UNIX_TIMESTAMP|CURRENT_DATE"), "UNIX_TIMESTAMP() and variants"},
		{regexp.MustCompile("(?i)\\(\\w+ OR \\w+\\) AS"), "OR in SELECT"},
		{regexp.MustCompile("(?i)SELECT \\(?[^)]+OR[^)]+\\)?.*FROM"), "OR in SELECT"},
		{regexp.MustCompile("(?i)^REPLACE (?:INTO)?.*VALUES"), "REPLACE INTO"},
		{regexp.MustCompile("(?i)^(?:RELEASE )?SAVEPOINT .+"), "SAVEPOINT/RELEASE SAVEPOINT"},
		{regexp.MustCompile("(?i)^ROLLBACK (?:WORK)?TO (?:SAVEPOINT)?.+"), "ROLLBACK SAVEPOINT"},
		{regexp.MustCompile("(?i)JOIN.*USING\\([^)]+\\)"), "JOIN ... USING(col)"},
		{regexp.MustCompile("(?i)CONVERT\\([^,)]+,[^,)]+\\)"), "CONVERT() function"},
		{regexp.MustCompile("(?i)SET (?:[^=]+=[^=,]+,)*[^=]+=\\s*(TRUE|FALSE)"), "BOOL as value"},
		{regexp.MustCompile("(?i)INSERT.*VALUES.*\\(([^),]+,\\s*)*(TRUE|FALSE)(\\s*,[^),]+)*\\)"), "BOOL as value"},
		{regexp.MustCompile("(?i)\\s*\\|\\|\\s*"), "|| in place of OR"},
		{regexp.MustCompile("(?i)ON DUPLICATE KEY UPDATE .*[^=]+=LAST_INSERT_ID"), "LAST_INSERT_ID in ON DUPLICATE KEY UPDATE"},
		{regexp.MustCompile("(?i)ORDER BY .* IS .+"), "IS in ORDER BY"},
		{regexp.MustCompile("(?i)GROUP_CONCAT\\([^)]+\\)"), "GROUP_CONCAT function"},
		{regexp.MustCompile("(?i)DIV \\s*\\d+"), "DIV operator"},
	}

	informationSchema = regexp.MustCompile("(?i)SELECT .+ FROM INFORMATION_SCHEMA")
	setVariables      = regexp.MustCompile("(?i)/\\*!\\d+ SET .*\\*/")

	flushLocalLogsPatterns = []*regexp.Regexp{
		regexp.MustCompile("(?i)^/usr/sbin/mysqld,\\s*Version:.*\\.\\s*started with:"),
		regexp.MustCompile("(?i)^Tcp port:\\s*\\d+\\s*Unix socket:.*"),
		regexp.MustCompile("(?i)^Time\\s*Id\\s*Command\\s*Argument"),
	}

	queryCache           = make(map[string]int)
	ignoreCache          = make(map[string]bool)
	unknownErrorExamples = make(map[string]string)
	handledErrorExamples = make(map[string]*handledErrorExample)

	unknownErrorCounts = make(map[string]int)
	handledErrorCounts = make(map[string]int)
	ignoredErrorCounts = make(map[string]map[string]int)

	linesSeen, linesMatched, linesWithoutQuery                        int
	unknownErrors, handledErrors, ignoredErrors                       int
	totalQueries, attemptedQueries, successfulQueries, skippedQueries int

	// todo: parse as both vttablet and vtgate?
)

func main() {
	defer exit.Recover()
	defer logutil.Flush()

	sigChan := make(chan os.Signal)
	go func() {
		stacktrace := make([]byte, 8192)
		for range sigChan {
			length := runtime.Stack(stacktrace, true)
			fmt.Println(string(stacktrace[:length]))
		}
	}()
	signal.Notify(sigChan, syscall.SIGQUIT)

	// flag parsing
	dbconfigFlags := dbconfigs.DbaConfig
	dbconfigs.RegisterFlags(dbconfigFlags)

	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, "Usage: %s [options] input_file \n", os.Args[0])
		fmt.Fprint(os.Stdout, "\nGiven an input_file general log, goes back tailBytes (default 10000) and attempts to parse queries in the log\n")
		fmt.Fprintf(os.Stdout, "\nOptions:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stdout, "\n")
	}

	flag.Parse()

	if len(flag.Args()) != 1 {
		flag.Usage()
		log.Errorf("vtparse requires an input_file positional argument")
		exit.Return(1)
	}
	servenv.FireRunHooks()

	ts := topo.Open()
	defer ts.Close()

	ctx, cancel := context.WithTimeout(context.Background(), *waitTime)
	wr := wrangler.New(logutil.NewConsoleLogger(), ts, tmclient.NewTabletManagerClient())

	installSignalHandlers(cancel)

	srv, err := wr.TopoServer().GetSrvVSchema(ctx, *cell)
	exitOnError(err, "get wrangler")
	vs, err := vindexes.BuildVSchema(srv)
	exitOnError(err, "build vschema")

	fileName := flag.Arg(0)
	stat, err := os.Stat(fileName)
	startSize := stat.Size()
	in, pos, err := openToLoc(fileName, startSize)

	exitOnError(err, "seek file %s", fileName)

	reader := bufio.NewReader(in)

	lineRe := regexp.MustCompile("(?:\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}\\.\\d{6}Z\\d*)(?:\\s+\\d+)?\\s+(\\w+)\\t*(?:(.+))?$")

	var buffer bytes.Buffer
	started := false

	creds := getCredentials()
	// simplify format string below by prepending :
	if val, ok := creds["password"]; ok && len(val) > 0 {
		creds["password"] = fmt.Sprintf(":%s", creds["password"])
	}

	var address string
	if val, ok := creds["socket"]; ok && len(val) > 0 {
		address = fmt.Sprintf("unix(%s)", creds["socket"])
	} else if val, ok := creds["host"]; ok && len(val) > 0 {
		address = fmt.Sprintf("tcp(%s)", creds["host"])
	} else {
		panic("Need to specify either socket or host, through -parse-conn-host, -parse-conn-socket, or -parse-conn-file")
	}
	db, err := sql.Open("mysql", fmt.Sprintf("%s%s@%s/%s", creds["user"], creds["password"], address, *keyspace))
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	tabletAlias, err := topoproto.ParseTabletAlias(*alias)
	exitOnError(err, "parse alias")

	tablet := &topodata.Tablet{
		Alias:          tabletAlias,
		Hostname:       *hostname,
		PortMap:        make(map[string]int32),
		Keyspace:       *keyspace,
		Shard:          "0",
		Type:           topodata.TabletType_REPLICA,
		DbNameOverride: *keyspace,
		Tags:           make(map[string]string),
	}

	err = wr.InitTablet(ctx, tablet, false, true, false)
	if err != nil && err != topo.ErrNodeExists {
		exitOnError(err, "init tablet")
	} else {
		defer cleanupTabletAndShard(ctx, wr, tabletAlias)
	}

	dbcfgs, err := dbconfigs.Init(creds["socket"], dbconfigFlags)
	exitOnError(err, "get dbconfigs")
	queryServiceStats := tabletserver.NewQueryServiceStats("vtparse", false)
	si := tabletserver.NewSchemaInfo("vtparse", dummyChecker{}, 0, 0, 0, make(map[string]string), false, queryServiceStats)
	cfg, err := dbconfigs.WithCredentials(&dbcfgs.Dba)
	si.Open(&cfg, false)
	err = si.Reload(ctx)
	exitOnError(err, "reload schema")

	err = topotools.RebuildVSchema(ctx, wr.Logger(), wr.TopoServer(), []string{*cell})
	exitOnError(err, "rebuild vschema")

	if *justInitialize {
		time.Sleep(10 * time.Second)
		fmt.Println("Initialized, exiting.")
		os.Exit(0)
	}
	out := os.Stdout
	if outFile != nil && len(*outFile) > 0 {
		dir := filepath.Dir(*outFile)
		os.MkdirAll(dir, 0755)
		out, err = os.OpenFile(*outFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		defer out.Close()
	}

	fmt.Fprintf(os.Stdout, "Reading from %d to %d\n", pos, startSize)

	var totalRead int64

	line, read, e := readLine(reader)

OUTER:
	for e == nil && pos <= startSize {
		pos += read
		totalRead += read
		linesSeen++

		matches := lineRe.FindStringSubmatch(line)
		if matches != nil {
			linesMatched++
			started = true
			if matches[1] != "Query" || strings.Contains(matches[2], "FLUSH LOCAL LOGS") {
				linesWithoutQuery++
				line, read, e = readLine(reader)
				continue
			}

			if buffer.Len() > 0 {
				totalQueries++
				data := string(buffer.Bytes())
				if ignored := ignored(data); ignored {
					if *verbosity > 1 {
						fmt.Fprintf(os.Stdout, "Skipped: %s\n", data)
					}
					skippedQueries++
				} else {
					attemptedQueries++
					if success := tryParse(data, vs, db); success {
						successfulQueries++
					}
				}
				buffer.Reset()
			}
			buffer.WriteString(matches[2])
		} else if started {
			for _, pattern := range flushLocalLogsPatterns {
				if pattern.MatchString(line) {
					line, read, e = readLine(reader)
					continue OUTER
				}
			}

			buffer.WriteString(" " + strings.TrimSpace(line) + " ")
		}

		line, read, e = readLine(reader)
	}

	if len(unknownErrorCounts) > 0 {
		out.WriteString("UNKNOWN ERRORS:\n")
		for errString, query := range unknownErrorExamples {
			count := unknownErrorCounts[errString]
			out.WriteString(fmt.Sprintf("\nQuery:\n%s\nError: %v\nCount: %d\n-----", query, errString, count))
		}
	}

	if len(handledErrorCounts) > 0 {
		out.WriteString("\n\nHANDLED ERRORS:\n")
		for _type, example := range handledErrorExamples {
			count := handledErrorCounts[_type]
			out.WriteString(fmt.Sprintf("\nQuery:\n%s\nError: %v\nType: %v\nCount: %d\n-----", example.sql, example.err, _type, count))
		}
	}

	fmt.Fprintf(os.Stdout, "Finished.\n\n"+
		"Lines read: %d\n"+
		"Lines matched: %d\n"+
		"Lines without queries: %d\n"+
		"Queries seen: %d\n"+
		"Queries attempted: %d\n"+
		"Queries succeeded: %d\n"+
		"Queries skipped: %d\n"+
		"Unknown errors: %d\n"+
		"Ignored errors: %d\n"+
		"Handled errors: %v\n"+
		"Ignored error breakdowns: %v\n"+
		"Handled error breakdowns: %v\n"+
		"Cache size: %d\n"+
		"Bytes read: %d\n"+
		"\n",
		linesSeen, linesMatched, linesWithoutQuery,
		totalQueries, attemptedQueries, successfulQueries, skippedQueries,
		unknownErrors, ignoredErrors, handledErrors, ignoredErrorCounts, handledErrorCounts, len(queryCache), totalRead)

	exit.Return(unknownErrors)
}

func cleanupTabletAndShard(ctx context.Context, wr *wrangler.Wrangler, tabletAlias *topodata.TabletAlias) {
	err := wr.DeleteTablet(ctx, tabletAlias, false)
	exitOnError(err, "delete tablet")
	err = wr.DeleteShard(ctx, *keyspace, "0", true, true)
	exitOnError(err, "delete shard")
}

func readLine(r *bufio.Reader) (string, int64, error) {
	var (
		isPrefix = true
		err      error
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), int64(len(ln)), err
}

func getCredentials() map[string]string {
	creds := make(map[string]string)
	creds["user"] = *connUser
	creds["password"] = *connPassword
	creds["host"] = *connHost
	creds["socket"] = *connSocket

	if connFile != nil && len(*connFile) > 0 {
		cin, cerr := os.Open(*connFile)
		exitOnError(cerr, "open conn file: %s", *connFile)
		cscanner := bufio.NewScanner(cin)

		re := regexp.MustCompile("(^[^=]+)=(.+)$")
		for cscanner.Scan() {
			line := cscanner.Text()
			matches := re.FindStringSubmatch(line)
			if matches != nil && len(matches) == 3 {
				creds[matches[1]] = matches[2]
			}
		}

	}

	// account for possibly quoted values
	re := regexp.MustCompile("'|\"")
	for key, value := range creds {
		if len(value) == 0 {
			continue
		}
		if re.MatchString(string(value[0])) && value[0] == value[len(value)-1] {
			creds[key] = value[1 : len(value)-1]

		}
	}

	// Default port
	re = regexp.MustCompile(":\\d+$")
	if !re.MatchString(creds["host"]) {
		creds["host"] = fmt.Sprintf("%s:3306", creds["host"])
	}

	return creds
}

func ignored(sql string) bool {
	if ignored, ok := ignoreCache[sql]; ok {
		return ignored
	}

	sql = strings.TrimSpace(strings.ToLower(sql))
	shouldIgnore := strings.HasPrefix(sql, "set ") ||
		strings.HasPrefix(sql, "/*!40103 set time_zone='+00:00' */") ||
		strings.HasPrefix(sql, "/*!40100 set @@sql_mode='' */") ||
		strings.HasPrefix(sql, "kill ") ||
		strings.HasPrefix(sql, "use ") ||
		strings.HasPrefix(sql, "show ") ||
		strings.HasPrefix(sql, "begin") ||
		strings.HasPrefix(sql, "commit") ||
		strings.HasPrefix(sql, "rollback") ||
		strings.HasPrefix(sql, "explain") ||
		strings.HasPrefix(sql, "flush") ||
		strings.HasPrefix(sql, "lock") ||
		strings.HasPrefix(sql, "unlock") ||
		strings.HasPrefix(sql, "purge") ||
		strings.HasPrefix(sql, "reset") ||
		strings.HasPrefix(sql, "change") ||
		strings.HasPrefix(sql, "start") ||
		strings.HasPrefix(sql, "stop") ||
		strings.HasPrefix(sql, "declare") ||
		strings.HasPrefix(sql, "desc") ||
		strings.HasPrefix(sql, "help") ||
		strings.HasPrefix(sql, "create") ||
		strings.HasPrefix(sql, "alter") ||
		strings.HasPrefix(sql, "drop") ||
		strings.HasPrefix(sql, "rename") ||
		strings.HasPrefix(sql, "truncate") ||
		strings.HasPrefix(sql, "load") ||
		strings.HasPrefix(sql, "grant") ||
		strings.HasPrefix(sql, "revoke") ||
		strings.HasPrefix(sql, "analyze") ||
		strings.HasPrefix(sql, "check") ||
		strings.HasPrefix(sql, "optimize") ||
		strings.HasPrefix(sql, "repair") ||
		strings.HasPrefix(sql, "install") ||
		strings.HasPrefix(sql, "uninstall") ||
		strings.Contains(sql, "memory_global_by_current_bytes") ||
		strings.Contains(sql, "heartbeat.heartbeat") ||
		strings.Contains(sql, "`heartbeat`.`heartbeat`") ||
		strings.Contains(sql, "`percona`.`checksums`") ||
		informationSchema.MatchString(sql) ||
		setVariables.MatchString(sql)

	ignoreCache[sql] = shouldIgnore

	return shouldIgnore
}

func exitOnError(err error, msg string, args ...interface{}) {
	if err != nil {
		if len(args) > 0 {
			msg = fmt.Sprintf(msg, args)
		}
		log.Errorf("Failed for: '%s', %v", msg, err)
		exit.Return(1)
	}
}

func openToLoc(fileName string, size int64) (*os.File, int64, error) {
	in, err := os.Open(fileName)
	var pos int64
	if *tailBytes < size {
		pos, err = in.Seek((*tailBytes)*-1, io.SeekEnd)
	}
	return in, pos, err
}

func tryParse(sql string, vs *vindexes.VSchema, db *sql.DB) bool {
	if res, ok := queryCache[sql]; ok {
		switch res {
		case 1:
			ignoredErrors++
		case 2:
			handledErrors++
		case 3:
			unknownErrors++
		}
		return res == 0
	}

	plan, err := planbuilder.Build(sql, &wrappedVSchema{
		vschema:  vs,
		keyspace: sqlparser.NewTableIdent(*keyspace),
	})

	if err != nil {
		if errIsIgnored(err) {
			if *verbosity > 1 {
				fmt.Fprintf(os.Stdout, "\nQuery:\n%s\nIgnored Error: %v\n-----", sql, err.Error())
			}
			ignoredErrors++
			queryCache[sql] = 1
		} else if _, ok := unknownErrorExamples[err.Error()]; ok {
			unknownErrorCounts[err.Error()]++
		} else {
			if errType, ok := errIsHandled(sql); ok {
				handledErrors++
				handledErrorCounts[errType]++
				if _, ok := handledErrorExamples[errType]; !ok {
					handledErrorExamples[errType] = &handledErrorExample{sql, err.Error()}
				}
				queryCache[sql] = 2
			} else {
				unknownErrors++
				unknownErrorCounts[err.Error()]++
				unknownErrorExamples[err.Error()] = sql
				queryCache[sql] = 3
			}
		}
		return false
	}

	route, _ := plan.Instructions.(*engine.Route)

	if *verbosity > 2 {
		fmt.Fprintf(os.Stdout, "Original: %s\n", sql)
		fmt.Fprintf(os.Stdout, "Parsed: %s\n", route.Query)
		fmt.Fprintf(os.Stdout, "Field: %s\n", route.FieldQuery)
	}

	if *compareExplain {
		res, err := db.Query(fmt.Sprintf("EXPLAIN %s", sql))
		if err == nil || !errIsIgnored(err) {
			source := parseRows(res, err)
			comp := parseRows(db.Query(fmt.Sprintf("EXPLAIN %s", route.Query)))

			if !reflect.DeepEqual(source, comp) {
				fmt.Fprintf(os.Stderr, "original explain: %v\n,parsed explain: %v", source, comp)
				panic("Explains not equal")
			}

			if len(route.FieldQuery) > 0 {
				fields := parseRows(db.Query(fmt.Sprintf("EXPLAIN %s", route.FieldQuery)))
				for _, field := range fields {
					if !field.extra.Valid || strings.ToLower(field.extra.String) != "impossible where" {
						fmt.Fprintf(os.Stderr, "field explain: %v", field)
						panic("FieldQuery not simple Impossible WHERE")
					}
				}
			} else if re := regexp.MustCompile("(?i)^\\s*INSERT|UPDATE|DELETE|SHOW|EXPLAIN|ALTER|DROP|CREATE"); !re.MatchString(sql) {
				panic("query should have a field query, but doesn't")
			}
		}
	}

	queryCache[sql] = 0
	return true

}

type explainResult struct {
	id           sql.NullInt64
	selectType   sql.NullString
	table        sql.NullString
	partitions   sql.NullString
	_type        sql.NullString
	possibleKeys sql.NullString
	key          sql.NullString
	keyLen       sql.NullString
	ref          sql.NullString
	rows         sql.NullInt64
	filtered     sql.NullString
	extra        sql.NullString
}

func parseRows(rows *sql.Rows, err error) []explainResult {
	exitOnError(err, "executing query")
	all := []explainResult{}

	count := 0
	for rows.Next() {
		parsed := explainResult{}
		err := rows.Scan(&parsed.id, &parsed.selectType, &parsed.table, &parsed.partitions, &parsed._type, &parsed.possibleKeys, &parsed.key, &parsed.keyLen, &parsed.ref, &parsed.rows, &parsed.filtered, &parsed.extra)
		count++
		exitOnError(err, "parsing row %d", count)
		all = append(all, parsed)

	}

	return all
}

func errIsIgnored(err error) bool {
	for _, pattern := range ignoredErrorPatterns {
		match := pattern.FindStringSubmatch(err.Error())
		if match != nil {
			if val, ok := ignoredErrorCounts[match[1]]; ok {
				val[match[2]]++
			} else {
				val := make(map[string]int)
				val[match[2]]++
				ignoredErrorCounts[match[1]] = val
			}
			return true
		}
	}

	return false
}

type handledError struct {
	pattern *regexp.Regexp
	errType string
}

type handledErrorExample struct {
	sql, err string
}

func errIsHandled(sql string) (string, bool) {
	for _, handled := range handledErrorPatterns {
		match := handled.pattern.MatchString(sql)
		if match {
			return handled.errType, true
		}
	}

	return "", false
}

type wrappedVSchema struct {
	vschema  *vindexes.VSchema
	keyspace sqlparser.TableIdent
}

func (vs *wrappedVSchema) Find(keyspace, tablename sqlparser.TableIdent) (table *vindexes.Table, err error) {
	if keyspace.IsEmpty() {
		keyspace = vs.keyspace
	}
	return vs.vschema.Find(keyspace.String(), tablename.String())
}

// signal handling, centralized here
func installSignalHandlers(cancel func()) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-sigChan
		// we got a signal, cancel the current ctx
		cancel()
	}()
}

type dummyChecker struct {
}

func (dummyChecker) CheckMySQL() {}
