// Copyright © 2015-2018 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package main

import (
	"path/filepath"
	"plugin"

	"github.com/platinasystems/go/goes/cmd"
	"github.com/platinasystems/go/goes/lang"
)

type vnetdCommand struct {
	plugin *plugin.Plugin
	cached struct {
		main     func(...string) error
		licenses func() map[string]string
		patents  func() map[string]string
		versions func() map[string]string
	}
}

var vnetd vnetdCommand

func (*vnetdCommand) String() string { return "vnetd" }

func (*vnetdCommand) Usage() string { return "vnetd" }

func (*vnetdCommand) Apropos() lang.Alt {
	return lang.Alt{
		lang.EnUS: "Platina's Mk1 TOR driver daemon",
	}
}

func (*vnetdCommand) Kind() cmd.Kind { return cmd.Daemon }

func (vnetd *vnetdCommand) Main(args ...string) error {
	if vnetd.cached.main == nil {
		vnetd.cached.main =
			vnetd.symbol("Main").(func(...string) error)
	}
	return vnetd.cached.main(args...)
}

func (vnetd *vnetdCommand) Licenses() map[string]string {
	if vnetd.cached.licenses == nil {
		vnetd.cached.licenses =
			vnetd.symbol("Licenses").(func() map[string]string)
	}
	return vnetd.cached.licenses()
}

func (vnetd *vnetdCommand) Patents() map[string]string {
	if vnetd.cached.patents == nil {
		vnetd.cached.patents =
			vnetd.symbol("Patents").(func() map[string]string)
	}
	return vnetd.cached.patents()
}

func (*vnetdCommand) Versions() map[string]string {
	if vnetd.cached.versions == nil {
		vnetd.cached.versions =
			vnetd.symbol("Versions").(func() map[string]string)
	}
	return vnetd.cached.versions()
}

func (vnetd *vnetdCommand) symbol(name string) plugin.Symbol {
	// FIXME first try unpacking zip file appended to Args[0]
	const so = "vnet-platina-mk1.so"
	libso := filepath.Join("/usr/lib/goes", so)
	if vnetd.plugin == nil {
		var err error
		vnetd.plugin, err = plugin.Open(libso)
		if err != nil {
			vnetd.plugin, err = plugin.Open(so)
			if err != nil {
				panic(err)
			}

		}
	}
	sym, err := vnetd.plugin.Lookup(name)
	if err != nil {
		panic(err)
	}
	return sym
}
