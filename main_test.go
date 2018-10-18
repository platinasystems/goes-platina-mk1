// Copyright © 2015-2018 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/platinasystems/fe1"
	"github.com/platinasystems/go/main/goes-platina-mk1/test"
	"github.com/platinasystems/go/platform/mk1"
	vnetFe1 "github.com/platinasystems/go/vnet/devices/ethernet/switch/fe1"
)

func Test(t *testing.T) {
	vnetFe1.AddPlatform = fe1.AddPlatform
	vnetFe1.Init = fe1.Init
	test.Suite(Machine, func() {
		if err := mk1.Start(Machine); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}, t)
}
