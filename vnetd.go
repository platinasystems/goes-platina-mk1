// Copyright © 2015-2018 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package main

import (
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/platinasystems/goes/cmd"
	"github.com/platinasystems/goes/lang"
)

type vnetd string

var Vnetd vnetd = "vnet-platina-mk1"

func (vnetd) String() string { return "vnetd" }

func (vnetd) Usage() string { return "vnetd" }

func (vnetd) Apropos() lang.Alt {
	return lang.Alt{
		lang.EnUS: "Platina's Mk1 TOR driver daemon",
	}
}

func (vnetd) Kind() cmd.Kind { return cmd.Daemon }

func (c vnetd) Main(args ...string) error {
	pn, err := c.Path()
	if err != nil {
		return err
	}
	if len(args) == 1 && strings.TrimLeft(args[0], "-") == "path" {
		fmt.Println(pn)
		return nil
	}
	return syscall.Exec(pn, append([]string{string(c)}, args...),
		os.Environ())
}

func (c vnetd) License() error {
	return c.Main("-license")
}

func (c vnetd) Patents() error {
	return c.Main("-patents")
}

func (c vnetd) Version() error {
	return c.Main("-version")
}

func (c vnetd) Path() (string, error) {
	cn := string(c)
	self, err := os.Readlink("/proc/self/exe")
	if err != nil {
		return "", err
	}
	if self == "/usr/bin/goes" {
		pn := "/usr/lib/goes/" + cn
		_, err = os.Stat(pn)
		if err == nil {
			return pn, nil
		}
		perr := err.(*os.PathError)
		return "", fmt.Errorf("%s: %s", perr.Path, perr.Err)
	}
	for _, pn := range []string{
		"./" + cn,
		"/usr/lib/goes/" + cn,
	} {
		if _, err := os.Stat(pn); err == nil {
			return pn, nil
		}
	}
	return "", fmt.Errorf("%s not found", cn)
}
