package main

import (
	"github.com/youtube/vitess/go/vt/topo"
)

// used at runtime by plug-ins
var (
	ts topo.Server
)