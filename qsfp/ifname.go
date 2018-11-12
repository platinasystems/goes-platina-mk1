// Copyright © 2016-2018 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qsfp

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"

	"github.com/platinasystems/redis"
)

const (
	numPorts           = 32
	alphaParameter     = "/sys/module/platina_mk1/parameters/alpha"
	provisionParameter = "/sys/module/platina_mk1/parameters/provision"
)

var cached struct {
	base, provision struct {
		once sync.Once
		val  interface{}
	}
}

func qsfpPortBase() int {
	cached.base.once.Do(func() {
		base := 1
		if f, err := os.Open(alphaParameter); err == nil {
			fmt.Fscan(f, &base)
			f.Close()
			if base > 1 {
				base = 1
			} else if base < 0 {
				base = 0
			}
		} else {
			s, err := redis.Hget(redis.DefaultHash, "eeprom.DeviceVersion")
			if err == nil {
				var ver int
				_, err = fmt.Sscan(s, &ver)
				if err == nil && (ver == 0 || ver == 0xff) {
					base = 0
				}
			}
		}
		cached.base.val = base
	})
	return cached.base.val.(int)
}

func qsfpProvision(i int) int {
	cached.provision.once.Do(func() {
		provision := make([]int, numPorts)
		buf, err := ioutil.ReadFile(provisionParameter)
		if err == nil {
			for i, s := range strings.Split(string(buf), ",") {
				if i < numPorts {
					fmt.Sscan(s, &provision[i])
				}
			}
		}
		cached.provision.val = provision
	})
	if i < numPorts {
		return cached.provision.val.([]int)[i]
	}
	return 0
}

// set 0 vs 1-base port numbering based on HW version
// use xethPORT-SUBPORT instead of xethPORT if provision[PORT] > 0
func qsfpIfnameOf(port, subport int) string {
	if qsfpProvision(port) == 0 {
		return fmt.Sprint("xeth", port+qsfpPortBase())
	}
	return fmt.Sprint("xeth", port+qsfpPortBase(), "-", subport+qsfpPortBase())
}
