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
)

var (
	tailBytes = flag.Int64("tailBytes", 10000, "bytes to parse of input file, starting at the end")
	waitTime = flag.Duration("wait-time", 1*time.Minute, "time to wait while parsing lines")
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

	srv, err := wr.TopoServer().GetSrvVSchema(ctx, "")
	exitOnError(err, "get wrangler")
	vs, err := vindexes.BuildVSchema(srv)
	exitOnError(err, "build vschema")

	fileName := flag.Arg(0)
	in, err := openToLoc(fileName)
	exitOnError(err, "seek file %s", fileName)

	scanner := bufio.NewScanner(in)

	re := regexp.MustCompile("(	?:\\d{4}-\\d{2}\\d{2}T\\d{2}:\\d{2}:\\d{2}:\\d{6}Z\\d*)(?:\\s+\\d+)?\\s+(\\w+)\\t*(?:(.+))?$")

	var buffer bytes.Buffer
	for scanner.Scan() {
		line := scanner.Text()

		matches := re.FindStringSubmatch(line)
		if matches != nil {
			if buffer.Len() > 0 {
				tryParse(string(buffer.Bytes()), vs)
				buffer.Reset()
			}
			buffer.WriteString(matches[0])
		} else {
			buffer.WriteString(strings.TrimSpace(line))
		}
	}

	log.Info("Finished")
	exit.Return(0)
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
	_, err = in.Seek(*tailBytes, 2)

	return in, err
}

func tryParse(sql string, vs *vindexes.VSchema) {
	plan, err  := planbuilder.Build(sql, &wrappedVSchema{
		vschema:  vs,
		keyspace: sqlparser.NewTableIdent(""),
	})
	exitOnError(err, "query=%v", sql)

	route, _ := plan.Instructions.(*engine.Route)
	log.Infof(route.FieldQuery)
	log.Infof(route.Query)
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
