// Copyright © 2016-2018 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"

	"github.com/platinasystems/goes/cmd"
	"github.com/platinasystems/goes/lang"
	"github.com/platinasystems/redis"
	"github.com/platinasystems/redis/publisher"
)

type tempdCommand chan struct{}

func (tempdCommand) String() string { return "tempd" }

func (tempdCommand) Usage() string { return "tempd" }

func (tempdCommand) Apropos() lang.Alt {
	return lang.Alt{
		lang.EnUS: "publish CPU temperature to redis",
	}
}

func (c tempdCommand) Close() error {
	close(c)
	return nil
}

func (tempdCommand) Kind() cmd.Kind { return cmd.Daemon }

func (c tempdCommand) Main(...string) error {
	var last string

	err := redis.IsReady()
	if err != nil {
		return err
	}

	pub, err := publisher.New()
	if err != nil {
		return err
	}

	t := time.NewTicker(5 * time.Second)
	defer t.Stop()
	for {
		select {
		case <-c:
			return nil
		case <-t.C:
			s := c.cpuCoreTemp()
			if s != last {
				pub.Print("sys.cpu.coretemp.C: ", s)
				last = s
			}
		}
	}

	return nil
}

func (c tempdCommand) cpuCoreTemp() string {
	var hi float64

	for _, x := range []struct {
		dir, dev string
	}{
		{"hwmon0", "core"}, // assumes lm-sensors
		{"hwmon1", "lm75"}, // assumes device is discovered
	} {
		t, err := c.readTemp(x.dir, x.dev)
		if err == nil && t > hi {
			hi = t
		}
	}
	s := fmt.Sprintf("%.2f\n", hi/1000)
	if s == "0" {
		s = ""
	}
	return s
}

func (c tempdCommand) readTemp(dir string, dev string) (float64, error) {
	const hwmon = "/sys/class/hwmon"
	var h float64
	n, err := ioutil.ReadFile(filepath.Join(hwmon, dir, "name"))
	if err != nil {
		return h, err
	}
	if !strings.HasPrefix(string(n), dev) {
		return h, nil
	}
	l, err := filepath.Glob(filepath.Join(hwmon, dir, "*_input"))
	if err != nil {
		return h, err
	}
	for _, fn := range l {
		t, err := ioutil.ReadFile(fn)
		if err == nil && len(t) > 0 {
			var tf float64
			fmt.Sscan(string(t), &tf)
			if tf > h {
				h = tf
			}
		}
	}
	return h, nil
}
