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
	"log"
	"os"

	"github.com/tangerine-network/tangerine-consensus/simulation"
	"github.com/tangerine-network/tangerine-consensus/simulation/config"
)

var configFile = flag.String("config", "", "path to simulation config file")

func main() {
	flag.Parse()

	if *configFile == "" {
		log.Println("error: no config file specified")
		os.Exit(1)
	}

	cfg, err := config.Read(*configFile)
	if err != nil {
		panic(err)
	}
	server := simulation.NewPeerServer()
	if _, err := server.Setup(cfg); err != nil {
		panic(err)
	}
	server.Run()
}
