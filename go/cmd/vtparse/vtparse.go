// Copyright 2017, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"time"
	log "github.com/golang/glog"
	"golang.org/x/net/context"
	"github.com/youtube/vitess/go/exit"
	"github.com/youtube/vitess/go/vt/dbconfigs"
	"github.com/youtube/vitess/go/vt/logutil"
	"fmt"
	"os"
	"github.com/youtube/vitess/go/vt/topo"
	"github.com/youtube/vitess/go/vt/wrangler"
	"github.com/youtube/vitess/go/vt/tabletmanager/tmclient"
	"github.com/youtube/vitess/go/vt/vtgate/planbuilder"
	"github.com/youtube/vitess/go/vt/sqlparser"
	"github.com/youtube/vitess/go/vt/vtgate/vindexes"
	"os/signal"
	"syscall"
	"github.com/youtube/vitess/go/vt/vtgate/engine"
	"bufio"
	"regexp"
	"bytes"
	"strings"
	"github.com/youtube/vitess/go/vt/servenv"
	"io"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
)

var (
	tailBytes = flag.Int64("tailBytes", 10000, "bytes to parse of input file, starting at the end")
	waitTime = flag.Duration("wait-time", 1*time.Minute, "time to wait while parsing lines")
	cell = flag.String("parse-cell", "", "cell to execute against")
	keyspace = flag.String("parse-keyspace", "", "keyspace to execute against")
	conn_user = flag.String("parse-conn-user", "root", "User to connect to db, i.e. root")
	conn_password = flag.String("parse-conn-password", "", "Password to connect to db")
	conn_host = flag.String("parse-conn-host", "localhost:3306", "Host and port to connect to db, i.e. localhost:3306")
	conn_socket = flag.String("parse-conn-socket", "", "Unix socket file to connect on, i.e. /var/run/mysql.sock")
	conn_file = flag.String("parse-conn-file", "", "Path to file containing connect info for db. Should contain at least one each of user=,host=,password=, on separate lines")

	ignored_error_patterns = []*regexp.Regexp {
		regexp.MustCompile("keyspace \\w+ not found in vschema"),
		regexp.MustCompile("symbol @@[\\w.]+ not found"),
	}
	query_cache = make(map[string]bool)
	ignore_cache = make(map[string]bool)
)

func main() {
	defer exit.Recover()
	defer logutil.Flush()

	// flag parsing
	dbconfigFlags := dbconfigs.DbaConfig
	dbconfigs.RegisterFlags(dbconfigFlags)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] input_file \n", os.Args[0])
		fmt.Fprint(os.Stderr, "\nGiven an input_file general log, goes back tailBytes (default 10000) and attempts to parse queries in the log\n")
		fmt.Fprintf(os.Stderr, "\nOptions:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\n")
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
	in, err := openToLoc(fileName)
	exitOnError(err, "seek file %s", fileName)

	scanner := bufio.NewScanner(in)

	line_re := regexp.MustCompile("(?:\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}\\.\\d{6}Z\\d*)(?:\\s+\\d+)?\\s+(\\w+)\\t*(?:(.+))?$")

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

	for scanner.Scan() {
		line := scanner.Text()

		matches := line_re.FindStringSubmatch(line)
		if matches != nil {
			started = true
			if matches[1] != "Query" {
				continue
			}

			if strings.Contains(matches[2], "FLUSH LOCAL LOGS") {
				scanner.Scan()
				scanner.Scan()
				scanner.Scan()
				continue
			}

			if buffer.Len() > 0 {
				sql := string(buffer.Bytes())
				if !ignored(sql) {
					tryParse(sql, vs, db)
				}
				buffer.Reset()
			}
			buffer.WriteString(matches[2])
		} else if started {
			buffer.WriteString(" " + strings.TrimSpace(line) + " ")
		}
	}
	exitOnError(scanner.Err(), "scanner error")
	fmt.Fprintf(os.Stderr, "Finished. Ignored %d queries, parsed %d\n", len(ignore_cache), len(query_cache))
	exit.Return(0)
}

func getCredentials() (map[string]string) {
	creds := make(map[string]string)
	creds["user"] = *conn_user
	creds["password"] = *conn_password
	creds["host"] = *conn_host
	creds["socket"] = *conn_socket

	if conn_file != nil && len(*conn_file) > 0 {
		cin, cerr := os.Open(*conn_file)
		exitOnError(cerr, "open conn file: %s", *conn_file)
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
		if re.MatchString(string(value[0])) && value[0] == value[len(value) - 1] {
			creds[key] = value[1:len(value)-1]

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
	if ignored, ok := ignore_cache[sql]; ok {
		return ignored
	}

	sql = strings.TrimSpace(strings.ToLower(sql))
	should_ignore := strings.HasPrefix(sql, "set ") ||
		strings.HasPrefix(sql, "use ") ||
		strings.HasPrefix(sql, "show ") ||
		strings.HasPrefix(sql, "begin") ||
		strings.HasPrefix(sql, "commit") ||
		strings.HasPrefix(sql, "explain") ||
		strings.Contains(sql, "memory_global_by_current_bytes") ||
		strings.Contains(sql, "heartbeat.heartbeat") ||
		strings.Contains(sql, "`heartbeat`.`heartbeat`")


	ignore_cache[sql] = should_ignore

	return should_ignore
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

func openToLoc(fileName string) (*os.File, error){
	in, err := os.Open(fileName)
	_, err = in.Seek((*tailBytes) * -1, io.SeekEnd)
	return in, err
}

func tryParse(sql string, vs *vindexes.VSchema, db *sql.DB) {
	if _, ok := query_cache[sql]; ok {
		return
	}

	plan, err := planbuilder.Build(sql, &wrappedVSchema{
		vschema:  vs,
		keyspace: sqlparser.NewTableIdent(*keyspace),
	})

	if err != nil && errIsIgnored(err) {
		query_cache[sql] = false
		return
	}

	exitOnError(err, "parse query=%v", sql)

	route, _ := plan.Instructions.(*engine.Route)
	fmt.Fprintf(os.Stderr, "Original: %s\n", sql)
	fmt.Fprintf(os.Stderr, "Parsed: %s\n", route.Query)
	fmt.Fprintf(os.Stderr, "Field: %s\n",route.FieldQuery)

	source := parseRows(db.Query(fmt.Sprintf("EXPLAIN %s", sql)))
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

	query_cache[sql] = true

}

type explainResult struct {
	id sql.NullInt64
	selectType sql.NullString
	table sql.NullString
	partitions  sql.NullString
	_type sql.NullString
	possible_keys sql.NullString
	key sql.NullString
	key_len sql.NullInt64
	ref sql.NullString
	rows sql.NullInt64
	filtered sql.NullString
	extra sql.NullString
}

func parseRows(rows *sql.Rows, err error) []explainResult {
	exitOnError(err, "executing query")
	all := []explainResult{}

	count := 0
	for rows.Next() {
		parsed := explainResult{}
		err := rows.Scan(&parsed.id, &parsed.selectType, &parsed.table, &parsed.partitions, &parsed._type, &parsed.possible_keys, &parsed.key, &parsed.key_len, &parsed.ref, &parsed.rows, &parsed.filtered, &parsed.extra)
		count++
		exitOnError(err, "parsing row %d", count)
		all = append(all, parsed)

	}

	return all
}

func errIsIgnored(err error) bool {
	if err == nil {
		return true
	}

	for _, pattern := range ignored_error_patterns {
		match := pattern.MatchString(err.Error())
		if match {
			return true
		}
	}

	return false
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
