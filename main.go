// Copyright © 2015-2018 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

// This is Platina's Mk1 TOR.
//
// To build this source,
//
//	go build
//
// Then zip and append an existing driver.
//
//	zip dirvers /usr/lib/goes/vnet-platina-mk1
//	cat drivers.zip >> goes-platina-mk1
//	zip -q -A goes-platina-mk1
//
// Or with NDA access to the driver source,
//
//	go build github.com/platinasystems/vnet-platina-mk1
//	zip drivers.zip vnet-platina-mk1
//
// Install the programs and driver with,
//
//	sudo ./goes-platina-mk1 install
package main

import (
	"fmt"
	"os"

	"github.com/platinasystems/redis"
)

func main() {
	redis.DefaultHash = "platina-mk1"
	if err := Goes.Main(os.Args...); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
