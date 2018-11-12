// Copyright © 2015-2018 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"

	"github.com/platinasystems/go/goes/cmd"
	"github.com/platinasystems/go/goes/lang"
)

type vnetdCommand struct{}

func (vnetdCommand) String() string { return "vnetd" }

func (vnetdCommand) Usage() string { return "vnetd" }

func (vnetdCommand) Apropos() lang.Alt {
	return lang.Alt{
		lang.EnUS: "Platina's Mk1 TOR driver daemon",
	}
}

func (vnetdCommand) Kind() cmd.Kind { return cmd.Daemon }

func (vnetdCommand) Main(args ...string) error {
	const vnetPlatinaMk1 = "vnet-platina-mk1"

	for _, dir := range []string{"/usr/lib/goes", "."} {
		fn := filepath.Join(dir, vnetPlatinaMk1)
		if _, err := os.Stat(fn); err == nil {
			return syscall.Exec(fn,
				append([]string{vnetPlatinaMk1}, args...),
				os.Environ())
		}
	}
	return fmt.Errorf("%s not found", vnetPlatinaMk1)
}

func (c vnetdCommand) License() error {
	return c.Main("-license")
}

func (c vnetdCommand) Patents() error {
	return c.Main("-patents")
}

func (c vnetdCommand) Version() error {
	return c.Main("-version")
}
