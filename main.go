// Copyright © 2015-2018 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

// This is Platina's Mk1 TOR.
//
// To build this source you'll first need to extract the driver plugin from an
// existing program binary.
//
//	unzip goes-platina-mk1 vnet-platina-mk1.so
//	zip plugins.zip vnet-platina-mk1.so
//
// Or with NDA access to the plugin source, build it with,
//
//	go build -buildmode=plugin github.com/platinasystems/vnet-platina-mk1
//	zip plugins.zip vnet-platina-mk1.so
//
// Then build the program and append the zipped plugin.
//
//	go build
//	cat plugins.zip >> goes-platina-mk1
//	zip -q -A goes-platina-mk1
//
// Install the programs and the plugin(s) with,
//
//	sudo ./goes-platina-mk1 install
package main

import (
	"fmt"
	"os"

	"github.com/platinasystems/go/goes"
	"github.com/platinasystems/redis"
)

func main() {
	goes.Info.Licenses = vnetd.Licenses
	goes.Info.Patents = vnetd.Patents
	goes.Info.Versions = func() map[string]string {
		m := vnetd.Versions()
		m["goes-platina-mk1"] = Version
		return m
	}
	redis.DefaultHash = "platina-mk1"
	if err := Goes.Main(os.Args...); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
