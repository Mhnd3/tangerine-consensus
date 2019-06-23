// Copyright 2018 The dexon-consensus Authors
// This file is part of the dexon-consensus library.
//
// The dexon-consensus library is free software: you can redistribute it
// and/or modify it under the terms of the GNU Lesser General Public License as
// published by the Free Software Foundation, either version 3 of the License,
// or (at your option) any later version.
//
// The dexon-consensus library is distributed in the hope that it will be
// useful, but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Lesser
// General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the dexon-consensus library. If not, see
// <http://www.gnu.org/licenses/>.

package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"
	"time"

	"gitlab.com/byzantine-lab/tangerine-consensus/simulation"
	"gitlab.com/byzantine-lab/tangerine-consensus/simulation/config"
)

var initialize = flag.Bool("init", false, "initialize config file")
var configFile = flag.String("config", "", "path to simulation config file")
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")
var logfile = flag.String("log", "", "write log to `file`-nodeID.log")

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	// Supports runtime pprof monitoring.
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	if *configFile == "" {
		fmt.Fprintln(os.Stderr, "error: no configuration file specified")
		os.Exit(1)
	}
	if *initialize {
		if err := config.GenerateDefault(*configFile); err != nil {
			fmt.Fprintf(os.Stderr, "error: %s", err)
			os.Exit(1)
		}
		//os.Exit(0)
	}

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	cfg, err := config.Read(*configFile)
	if err != nil {
		panic(err)
	}
	simulation.Run(cfg, *logfile)

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		// #nosec G104
		f.Close()
	}
}
