// Copyright © 2015-2018 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/platinasystems/go/goes"
	"github.com/platinasystems/go/main/goes-platina-mk1/test"
	"github.com/platinasystems/go/platform/mk1"
	"github.com/platinasystems/go/vnet/devices/ethernet/switch/fe1"
)

const Machine = "goes-platina-mk1"

func Test(t *testing.T) {
	test.Suite(Machine, func() {
		goes.Info.Licenses = Licenses
		goes.Info.Patents = Patents
		goes.Info.Versions = Versions
		fe1.AddPlatform = AddPlatform
		fe1.Init = Init
		if err := mk1.Start(Machine); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}, t)
}
