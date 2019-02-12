// Copyright © 2017-2019 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/platinasystems/goes/lang"
)

type status struct{}

func (status) Apropos() lang.Alt {
	return lang.Alt{
		lang.EnUS: "print status of goes-platina-mk1",
	}
}

func (status) String() string { return "status" }
func (status) Usage() string  { return "[show ]status" }

func (status status) Main(args ...string) error {
	if os.Geteuid() != 0 {
		return errors.New("you aren't root")
	}
	if len(args) > 0 {
		return fmt.Errorf("%v: unexpected", args)
	}
	goes, err := os.Readlink("/proc/self/exe")
	if err != nil {
		return err
	}
	fmt.Println("GOES status")
	fmt.Println("======================")
	fmt.Printf("  %-15s - %s\n", "Mode", "XETH")

	for _, x := range []struct {
		header string
		f      func(string) error
	}{
		{"PCI", status.chip},
		{"Check daemons", status.daemons},
		{"Check Redis", status.redis},
		{"Check vnet", status.vnet},
	} {
		fmt.Printf("  %-15s - ", x.header)
		if err := x.f(goes); err == nil {
			fmt.Println("OK")
		} else {
			fmt.Printf("Not OK\n")
			return err
		}
	}

	return nil
}

func (status) chip(goes string) error {
	const thpat = "Broadcom (Corporation|Limited) Device b96[05]"

	out, err := exec.Command("/usr/bin/lspci").Output()
	if err != nil {
		return err
	}

	match, err := regexp.Match(thpat, out)
	if err != nil {
		return err
	}

	if !match {
		err = fmt.Errorf("TH missing")
	}
	return err
}

func (status) daemons(goes string) error {
	out, err := exec.Command(goes, "show", "daemons").Output()
	if err != nil {
		return err
	}
	sout := string(out)
	for _, daemon := range []string{
		"[redisd]",
		"[uptimed]",
		"[tempd]",
		"[vnetd]",
	} {
		if strings.Index(sout, daemon) < 0 {
			return fmt.Errorf("missing %s from:\n%s", daemon, sout)
		}
	}
	return nil
}

func (status) redis(goes string) error {
	out, err := exec.Command(goes,
		"hget", "platina-mk1", "redis.ready").Output()
	if err != nil {
		fmt.Println(err, out)
		err = fmt.Errorf("%v\n%s", err, string(out))
	}
	return err
}

func (status) vnet(goes string) error {
	args := []string{"/usr/bin/timeout", "30", goes,
		"vnet", "show", "hardware"}
	_, err := exec.Command(args[0], args[1:]...).Output()
	if err != nil {
		return fmt.Errorf("vnetd daemon not responding")
	}
	return nil
}
