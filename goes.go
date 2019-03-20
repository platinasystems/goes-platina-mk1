// Copyright © 2015-2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/platinasystems/goes"
	"github.com/platinasystems/goes/cmd"
	"github.com/platinasystems/goes/cmd/bang"
	"github.com/platinasystems/goes/cmd/biosupdate"
	"github.com/platinasystems/goes/cmd/buildid"
	"github.com/platinasystems/goes/cmd/cat"
	"github.com/platinasystems/goes/cmd/cd"
	"github.com/platinasystems/goes/cmd/chmod"
	"github.com/platinasystems/goes/cmd/cli"
	"github.com/platinasystems/goes/cmd/cmdline"
	"github.com/platinasystems/goes/cmd/cp"
	"github.com/platinasystems/goes/cmd/daemons"
	"github.com/platinasystems/goes/cmd/dmesg"
	"github.com/platinasystems/goes/cmd/echo"
	eeprom "github.com/platinasystems/goes/cmd/eeprom/platina_eeprom"
	"github.com/platinasystems/goes/cmd/elsecmd"
	"github.com/platinasystems/goes/cmd/env"
	"github.com/platinasystems/goes/cmd/exec"
	"github.com/platinasystems/goes/cmd/exit"
	"github.com/platinasystems/goes/cmd/export"
	"github.com/platinasystems/goes/cmd/falsecmd"
	"github.com/platinasystems/goes/cmd/femtocom"
	"github.com/platinasystems/goes/cmd/ficmd"
	"github.com/platinasystems/goes/cmd/function"
	"github.com/platinasystems/goes/cmd/gpio"
	"github.com/platinasystems/goes/cmd/hdel"
	"github.com/platinasystems/goes/cmd/hdelta"
	"github.com/platinasystems/goes/cmd/hexists"
	"github.com/platinasystems/goes/cmd/hget"
	"github.com/platinasystems/goes/cmd/hgetall"
	"github.com/platinasystems/goes/cmd/hkeys"
	"github.com/platinasystems/goes/cmd/hset"
	"github.com/platinasystems/goes/cmd/hwait"
	"github.com/platinasystems/goes/cmd/ifcmd"
	"github.com/platinasystems/goes/cmd/iminfo"
	"github.com/platinasystems/goes/cmd/insmod"
	"github.com/platinasystems/goes/cmd/install"
	"github.com/platinasystems/goes/cmd/iocmd"
	"github.com/platinasystems/goes/cmd/ip"
	"github.com/platinasystems/goes/cmd/kexec"
	"github.com/platinasystems/goes/cmd/keys"
	"github.com/platinasystems/goes/cmd/kill"
	"github.com/platinasystems/goes/cmd/ln"
	"github.com/platinasystems/goes/cmd/log"
	"github.com/platinasystems/goes/cmd/ls"
	"github.com/platinasystems/goes/cmd/lsmod"
	"github.com/platinasystems/goes/cmd/mac_ll"
	"github.com/platinasystems/goes/cmd/mkdir"
	"github.com/platinasystems/goes/cmd/mknod"
	"github.com/platinasystems/goes/cmd/mount"
	"github.com/platinasystems/goes/cmd/ping"
	"github.com/platinasystems/goes/cmd/platina/mk1/bootc"
	"github.com/platinasystems/goes/cmd/platina/mk1/toggle"
	"github.com/platinasystems/goes/cmd/platina/mk1/upgrade"
	"github.com/platinasystems/goes/cmd/ps"
	"github.com/platinasystems/goes/cmd/pwd"
	"github.com/platinasystems/goes/cmd/reboot"
	"github.com/platinasystems/goes/cmd/redisd"
	"github.com/platinasystems/goes/cmd/reload"
	"github.com/platinasystems/goes/cmd/restart"
	"github.com/platinasystems/goes/cmd/rm"
	"github.com/platinasystems/goes/cmd/rmmod"
	"github.com/platinasystems/goes/cmd/slashinit"
	"github.com/platinasystems/goes/cmd/sleep"
	"github.com/platinasystems/goes/cmd/source"
	"github.com/platinasystems/goes/cmd/start"
	"github.com/platinasystems/goes/cmd/stop"
	"github.com/platinasystems/goes/cmd/stty"
	"github.com/platinasystems/goes/cmd/subscribe"
	"github.com/platinasystems/goes/cmd/sync"
	"github.com/platinasystems/goes/cmd/testcmd"
	"github.com/platinasystems/goes/cmd/thencmd"
	"github.com/platinasystems/goes/cmd/truecmd"
	"github.com/platinasystems/goes/cmd/umount"
	"github.com/platinasystems/goes/cmd/uninstall"
	"github.com/platinasystems/goes/cmd/uptimed"
	"github.com/platinasystems/goes/cmd/wget"
	"github.com/platinasystems/goes/lang"
	"github.com/platinasystems/redis"
	"github.com/platinasystems/redis/publisher"
)

const (
	onieDir = "/sys/bus/i2c/devices/0-0051/onie"
)

var onieFileToRedisName = map[string]string{
	"country_code":     "CountryCode",
	"crc":              "Crc",
	"device_version":   "DeviceVersion",
	"diag_version":     "DiagVersion",
	"label_revision":   "LabelRevision",
	"mac_base":         "BaseEthernetAddress",
	"manufacture_date": "ManufactureDate",
	"manufacturer":     "Manufacturer",
	"num_macs":         "NEthernetAddress",
	"onie_version":     "OnieVersion",
	"part_number":      "PartNumber",
	"platform_name":    "PlatformName",
	"product_name":     "ProductName",
	"serial_number":    "SerialNumber",
	"service_tag":      "ServiceTag",
	"vendor":           "Vendor",
	"vendor_extension": "VendorExtension",
}

var Goes = &goes.Goes{
	NAME: "goes-platina-mk1",
	APROPOS: lang.Alt{
		lang.EnUS: "goes machine for platina's mk1 TOR",
	},
	ByName: map[string]cmd.Cmd{
		"!":          bang.Command{},
		"cli":        &cli.Command{},
		"biosupdate": biosupdate.Command{},
		"bootc":      &bootc.Command{},
		"buildid":    buildid.Command{},
		"cat":        cat.Command{},
		"cd":         &cd.Command{},
		"chmod":      chmod.Command{},
		"copyright":  License,
		"cp":         cp.Command{},
		"daemons":    daemons.Admin,
		"dmesg":      dmesg.Command{},
		"echo":       echo.Command{},
		"else":       &elsecmd.Command{},
		"env":        &env.Command{},
		"exec":       exec.Command{},
		"exit":       exit.Command{},
		"export":     export.Command{},
		"false":      falsecmd.Command{},
		"femtocom":   femtocom.Command{},
		"fi":         &ficmd.Command{},
		"function":   &function.Command{},
		"gpio":       &gpio.Command{},
		"goes-daemons": &daemons.Server{
			Init: [][]string{
				[]string{"redisd"},
				[]string{"uptimed"},
				[]string{"tempd"},
				[]string{"vnetd"},
			},
		},
		"hdel":    hdel.Command{},
		"hdelta":  &hdelta.Command{},
		"hexists": hexists.Command{},
		"hget":    hget.Command{},
		"hgetall": hgetall.Command{},
		"hkeys":   hkeys.Command{},
		"hset":    hset.Command{},
		"hwait":   hwait.Command{},
		"if":      &ifcmd.Command{},
		"insmod":  insmod.Command{},
		"install": &install.Command{},
		"io":      iocmd.Command{},
		"ip":      ip.Goes,
		"kexec":   kexec.Command{},
		"keys":    keys.Command{},
		"kill":    kill.Command{},
		"license": License,
		"ln":      ln.Command{},
		"log":     log.Command{},
		"ls":      ls.Command{},
		"lsmod":   lsmod.Command{},
		"mkdir":   mkdir.Command{},
		"mknod":   mknod.Command{},
		"mac-ll":  mac_ll.Command{},
		"mount":   mount.Command{},
		"patents": Patents,
		"ping":    ping.Command{},
		"ps":      ps.Command{},
		"pwd":     pwd.Command{},
		"reboot":  &reboot.Command{},
		"redisd": &redisd.Command{
			Devs:    []string{"lo", "eth0"},
			Machine: string(Machine),
			Hook: func(pub *publisher.Publisher) {
				_, err := os.Stat(onieDir)
				if err == nil {
					oniePub(pub)
				} else {
					platinaMk1EepromConfig()
					eeprom.RedisdHook(pub)
				}
			},
		},
		"reload":  reload.Command{},
		"restart": &restart.Command{},
		"rm":      rm.Command{},
		"rmmod":   rmmod.Command{},
		"show": &goes.Goes{
			NAME:  "show",
			USAGE: "show OBJECT",
			APROPOS: lang.Alt{
				lang.EnUS: "print stuff",
			},
			ByName: map[string]cmd.Cmd{
				"buildid":   buildid.Command{},
				"cmdline":   cmdline.Command{},
				"copyright": License,
				"daemons":   daemons.Status{},
				"iminfo":    iminfo.Command{},
				"license":   License,
				"log":       daemons.Log{},
				"machine":   Machine,
				"patents":   Patents,
				"status":    status{},
				"tech":      tech{},
				"version":   &Version,
			},
		},
		"/init":  &slashinit.Command{},
		"sleep":  sleep.Command{},
		"source": &source.Command{},
		"start": &start.Command{
			ConfHook: func() error {
				return redis.Hwait(redis.DefaultHash,
					"vnet.ready", "true",
					10*time.Second)
			},
		},
		"stop":      &stop.Command{},
		"status":    status{},
		"stty":      stty.Command{},
		"subscribe": subscribe.Command{},
		"sync":      sync.Command{},
		"tempd":     tempdCommand(make(chan struct{})),
		"[":         testcmd.Command{},
		"then":      &thencmd.Command{},
		"toggle":    &toggle.Command{},
		"true":      truecmd.Command{},
		"umount":    umount.Command{},
		"uninstall": &uninstall.Command{},
		"upgrade":   upgrade.Command{},
		"uptimed":   uptimed.Command(make(chan struct{})),
		"-version":  &Version,
		"--version": &Version,
		"version":   &Version,
		"vnet":      vnetCommand{},
		"vnetd":     Vnetd,
		"wget":      wget.Command{},
	},
}

func platinaMk1EepromConfig() {
	eeprom.Config(
		eeprom.BusIndex(0),
		eeprom.BusAddress(0x51),
		eeprom.BusDelay(10*time.Millisecond),
		eeprom.MinMacs(132),
		eeprom.OUI([3]byte{0x02, 0x46, 0x8a}),
	)
}

func oniePub(pub *publisher.Publisher) {
	list, err := ioutil.ReadDir(onieDir)
	if err != nil {
		pub.Printf("onie: %v", err)
		return
	}
	for _, fi := range list {
		var s string
		data, err := ioutil.ReadFile(filepath.Join(onieDir, fi.Name()))
		if err == nil {
			s = string(data)
		} else {
			s = err.Error()
		}
		pub.Printf("eeprom.%s: %s", onieFileToRedisName[fi.Name()], s)
	}
}
