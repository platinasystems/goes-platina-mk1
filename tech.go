// Copyright © 2019 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/platinasystems/goes/lang"
)

type tech struct{}

func (tech) String() string { return "tech" }

func (tech) Usage() string { return "show tech[ support]" }

func (tech) Apropos() lang.Alt {
	return lang.Alt{
		lang.EnUS: "Print info for Platina technician",
	}
}

type techEntry struct {
	// nested key depth
	level int
	key   string
	// show conditionally
	skip bool
	// if block is true, add "|" style to key and indent program output
	block bool
	// if nested the prog output is indented further
	nested bool
	// if prog is empty, just print space separated args
	prog string
	args []string
}

func (c tech) Main(args ...string) error {
	var (
		skipOnie   = true
		skipRedis  = true
		skipVnet   = true
		skipStatus = os.Geteuid() != 0
	)
	_, err := os.Stat("/sys/bus/i2c/devices/0-0051/onie")
	if err == nil {
		skipOnie = false
	}
	data, err := ioutil.ReadFile("/proc/net/unix")
	if err == nil {
		if bytes.Index(data, []byte("@vnet\n")) > 0 {
			skipVnet = false
		}
		if bytes.Index(data, []byte("@redisd\n")) > 0 {
			skipRedis = false
		}
	}

	w := io.Writer(os.Stdout)

	self, err := os.Readlink("/proc/self/exe")
	if err != nil {
		return err
	}
	vnetd, err := Vnetd.Path()
	if err != nil {
		return err
	}

	fmt.Fprintln(w, "---")
	for _, entry := range []techEntry{
		{
			key:    "eeprom",
			skip:   skipOnie,
			nested: true,
			prog:   self,
			args:   []string{"show", "onie"},
		},
		{
			skip:   !skipOnie || skipRedis,
			nested: true,
			prog:   self,
			args:   []string{"hget", "platina-mk1", "eeprom"},
		},
		{
			key: filepath.Base(self),
		},
		{
			level: 1,
			key:   "path",
			args:  []string{self},
		},
		{
			level: 1,
			key:   "buildid",
			prog:  self,
			args:  []string{"show", "buildid"},
		},
		{
			level: 1,
			key:   "version",
			args:  []string{string(Version)},
		},
		{
			key: string(Vnetd),
		},
		{
			level: 1,
			key:   "path",
			args:  []string{vnetd},
		},
		{
			level: 1,
			key:   "buildid",
			prog:  self,
			args:  []string{"show", "buildid", vnetd},
		},
		{
			level:  1,
			key:    "version",
			nested: true,
			prog:   vnetd,
			args:   []string{"-version"},
		},
		{
			key:    "platina-mk1.ko",
			nested: true,
			prog:   "/sbin/modinfo",
			args:   []string{"platina-mk1"},
		},
		{
			key:   "status",
			skip:  skipStatus,
			block: true,
			prog:  self,
			args:  []string{"status"},
		},
		{
			key:   "/dev/kmsg",
			block: true,
			prog:  "/bin/dmesg",
			args:  []string{"-H"},
		},
		{
			key:   "log",
			block: true,
			prog:  self,
			args:  []string{"show", "log"},
		},
		{
			key:   "vnet_errors",
			skip:  skipVnet,
			block: true,
			prog:  self,
			args:  []string{"vnet", "show", "errors"},
		},
		{
			key:   "vnet_hardware",
			skip:  skipVnet,
			block: true,
			prog:  self,
			args:  []string{"vnet", "show", "hardware"},
		},
		{
			key:   "vnet_fe1_int",
			skip:  skipVnet,
			block: true,
			prog:  self,
			args:  []string{"vnet", "show", "fe1", "int"},
		},
		{
			key:   "vnet_fe1_pipe",
			skip:  skipVnet,
			block: true,
			prog:  self,
			args:  []string{"vnet", "show", "fe1", "pipe"},
		},
		{
			key:   "vnet_fe1_switches",
			skip:  skipVnet,
			block: true,
			prog:  self,
			args:  []string{"vnet", "show", "fe1", "switches"},
		},
		{
			key:   "vnet_fe1_temp",
			skip:  skipVnet,
			block: true,
			prog:  self,
			args:  []string{"vnet", "show", "fe1", "temp"},
		},
		{
			key:   "vnet_fe1_port_phy",
			skip:  skipVnet,
			block: true,
			prog:  self,
			args:  []string{"vnet", "show", "fe1", "port", "phy"},
		},
		{
			key:   "vnet_fe1_serdes",
			skip:  skipVnet,
			block: true,
			prog:  self,
			args:  []string{"vnet", "show", "fe1", "serdes"},
		},
		{
			key:   "vnet_fe1_port_map_vlan",
			skip:  skipVnet,
			block: true,
			prog:  self,
			args:  []string{"vnet", "show", "fe1", "port-map", "vlan"},
		},
		{
			key:   "vnet_fe1_port_tab",
			skip:  skipVnet,
			block: true,
			prog:  self,
			args:  []string{"vnet", "show", "fe1", "port-tab"},
		},
		{
			key:   "vnet_fe1_vlan",
			skip:  skipVnet,
			block: true,
			prog:  self,
			args:  []string{"vnet", "show", "fe1", "vlan"},
		},
		{
			key:   "vnet_fe1_acl_l2",
			skip:  skipVnet,
			block: true,
			prog:  self,
			args:  []string{"vnet", "show", "fe1", "acl", "l2"},
		},
		{
			key:   "vnet_fe1_acl_l3",
			skip:  skipVnet,
			block: true,
			prog:  self,
			args:  []string{"vnet", "show", "fe1", "acl", "l3"},
		},
		{
			key:   "vnet_fe1_l2",
			skip:  skipVnet,
			block: true,
			prog:  self,
			args:  []string{"vnet", "show", "fe1", "l2"},
		},
		{
			key:   "vnet_fe1_station",
			skip:  skipVnet,
			block: true,
			prog:  self,
			args:  []string{"vnet", "show", "fe1", "station"},
		},
		{
			key:   "vnet_fe1_vlan_br",
			skip:  skipVnet,
			block: true,
			prog:  self,
			args:  []string{"vnet", "show", "fe1", "vlan", "br"},
		},
		{
			key:   "vnet_fe1_vlan_rx",
			skip:  skipVnet,
			block: true,
			prog:  self,
			args:  []string{"vnet", "show", "fe1", "vlan", "rx"},
		},
		{
			key:   "vnet_fe1_vlan_tx",
			skip:  skipVnet,
			block: true,
			prog:  self,
			args:  []string{"vnet", "show", "fe1", "vlan", "tx"},
		},
		{
			key:   "vnet_fe1_tcam",
			skip:  skipVnet,
			block: true,
			prog:  self,
			args:  []string{"vnet", "show", "fe1", "tcam"},
		},
		{
			key:   "vnet_fe1_l3_rx",
			skip:  skipVnet,
			block: true,
			prog:  self,
			args:  []string{"vnet", "show", "fe1", "l3", "rx"},
		},
		{
			key:   "vnet_fe1_l3_tx",
			skip:  skipVnet,
			block: true,
			prog:  self,
			args:  []string{"vnet", "show", "fe1", "l3", "tx"},
		},
		{
			key:   "vnet_ports",
			skip:  skipVnet,
			block: true,
			prog:  self,
			args:  []string{"vnet", "show", "ports"},
		},
		{
			key:   "vnet_int",
			skip:  skipVnet,
			block: true,
			prog:  self,
			args:  []string{"vnet", "show", "int"},
		},
		{
			key:   "vnet_buf",
			skip:  skipVnet,
			block: true,
			prog:  self,
			args:  []string{"vnet", "show", "buf"},
		},
		{
			key:   "vnet_ip_fib",
			skip:  skipVnet,
			block: true,
			prog:  self,
			args:  []string{"vnet", "show", "ip", "fib"},
		},
		{
			key:   "vnet_neighbor",
			skip:  skipVnet,
			block: true,
			prog:  self,
			args:  []string{"vnet", "show", "neighbor"},
		},
		{
			key:   "vnet_br",
			skip:  skipVnet,
			block: true,
			prog:  self,
			args:  []string{"vnet", "show", "br"},
		},
	} {
		if !entry.skip {
			c.show(w, &entry)
		}
	}
	fmt.Fprintln(w, "...")
	return err
}

func (tech) show(w io.Writer, entry *techEntry) {
	indent := func() {}
	if len(entry.key) > 0 {
		for i := 0; i < entry.level; i++ {
			fmt.Fprint(w, "  ")
		}
		fmt.Fprint(w, entry.key, ": ")
		if entry.nested {
			fmt.Fprintln(w)
			indent = func() {
				for i := 0; i <= entry.level; i++ {
					fmt.Fprint(w, "  ")
				}
			}
		} else if entry.block {
			fmt.Fprintln(w, "|")
			indent = func() {
				fmt.Fprint(w, "    ")
			}
		}
	}
	if len(entry.prog) == 0 {
		for i, arg := range entry.args {
			if i > 0 {
				fmt.Fprint(w, " ")
			}
			fmt.Fprint(w, arg)
		}
		fmt.Fprintln(w)
		return
	}
	n := 0
	cmd := exec.Command(entry.prog, entry.args...)
	pout, err := cmd.StdoutPipe()
	if err == nil {
		if err = cmd.Start(); err == nil {
			scanner := bufio.NewScanner(pout)
			for scanner.Scan() {
				t := scanner.Text()
				indent()
				if entry.block {
					fmt.Fprintln(w, t)
				} else if entry.nested {
					fields := strings.Fields(t)
					if len(fields) == 1 {
						t = fields[0] + " null"
					} else {
						t = strings.Join(fields, " ")
					}
					fmt.Fprintln(w, t)
				} else if n == 0 {
					fmt.Fprintln(w, t)
				}
				n += 1
			}
			err = cmd.Wait()
		}
	}
	if n == 0 {
		if err != nil {
			indent()
			if entry.nested {
				fmt.Fprintln(w, "error:", err.Error())
			} else {
				fmt.Fprintln(w, err.Error())
			}
		} else {
			fmt.Fprintln(w)
		}
	}
}
