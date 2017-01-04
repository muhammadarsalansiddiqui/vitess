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
)

var (
	tailBytes = flag.Int64("tailBytes", 10000, "bytes to parse of input file, starting at the end")
	waitTime = flag.Duration("wait-time", 1*time.Minute, "time to wait while parsing lines")
	cell = flag.String("cell", "", "cell to execute against")
	keyspace = flag.String("keyspace", "", "keyspace to execute against")
	ignored_error_patterns = []*regexp.Regexp {
		regexp.MustCompile("keyspace \\w+ not found in vschema"),
		regexp.MustCompile("symbol @@[\\w.]+ not found"),
	}
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

	re := regexp.MustCompile("(?:\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}\\.\\d{6}Z\\d*)(?:\\s+\\d+)?\\s+(\\w+)\\t*(?:(.+))?$")

	var buffer bytes.Buffer
	started := false

	for scanner.Scan() {
		line := scanner.Text()

		matches := re.FindStringSubmatch(line)
		if matches != nil {
			started = true
			if matches[1] != "Query" {
				continue
			}

			if buffer.Len() > 0 {
				sql := string(buffer.Bytes())
				if !ignored(sql) {
					fmt.Fprintf(os.Stderr, "Trying to parse: %s\n", sql)
					tryParse(sql, vs)
				}
				buffer.Reset()
			}
			buffer.WriteString(matches[2])
		} else if started {
			fmt.Fprintln(os.Stderr, "No match\n")
			buffer.WriteString(" " + strings.TrimSpace(line) + " ")
		}
	}
	exitOnError(scanner.Err(), "scanner error")
	fmt.Fprintln(os.Stderr, "Finished")
	exit.Return(0)
}

func ignored(sql string) bool {
	sql = strings.ToLower(sql)
	return strings.HasPrefix(sql, "set ") ||
		strings.HasPrefix(sql, "use ") ||
		strings.HasPrefix(sql, "show ") ||
		strings.HasPrefix(sql, "begin") ||
		strings.HasPrefix(sql, "commit") ||
		strings.Contains(sql, "memory_global_by_current_bytes") ||
		strings.Contains(sql, "heartbeat.heartbeat") ||
		strings.Contains(sql, "`heartbeat`.`heartbeat`")
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

func tryParse(sql string, vs *vindexes.VSchema) {
	plan, err := planbuilder.Build(sql, &wrappedVSchema{
		vschema:  vs,
		keyspace: sqlparser.NewTableIdent(*keyspace),
	})

	if err != nil && errIsIgnored(err) {
		return
	}

	exitOnError(err, "query=%v", sql)

	route, _ := plan.Instructions.(*engine.Route)
	fmt.Fprintf(os.Stderr, "Original: %s\n", sql)
	fmt.Fprintf(os.Stderr, "Parsed: %s\n", route.Query)
	fmt.Fprintf(os.Stderr, "Field: %s\n",route.FieldQuery)

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
