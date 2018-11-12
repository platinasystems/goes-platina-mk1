// Copyright © 2015-2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package qsfp

import (
	"fmt"
	"strconv"

	"github.com/platinasystems/redis"
)

func qsfpInit() {
	var ver, portOffset int
	s, err := redis.Hget(redis.DefaultHash, "eeprom.DeviceVersion")
	if err == nil {
		_, err = fmt.Sscan(s, &ver)
		if err == nil && (ver == 0 || ver == 0xff) {
			portOffset = -1
		}
	}

	//port 1-16 present signals
	qsfpVdevIo[0] = qsfpI2cDev{0, 0x20, 0, 0x70, 0x10, 0, 0, 0}
	//port 17-32 present signals
	qsfpVdevIo[1] = qsfpI2cDev{0, 0x21, 0, 0x70, 0x10, 0, 0, 0}
	//port 1-16 interrupt signals
	qsfpVdevIo[2] = qsfpI2cDev{0, 0x22, 0, 0x70, 0x10, 0, 0, 0}
	//port 17-32 interrupt signals
	qsfpVdevIo[3] = qsfpI2cDev{0, 0x23, 0, 0x70, 0x10, 0, 0, 0}
	//port 1-16 LP mode signals
	qsfpVdevIo[4] = qsfpI2cDev{0, 0x20, 0, 0x70, 0x20, 0, 0, 0}
	//port 17-32 LP mode signals
	qsfpVdevIo[5] = qsfpI2cDev{0, 0x21, 0, 0x70, 0x20, 0, 0, 0}
	//port 1-16 reset signals
	qsfpVdevIo[6] = qsfpI2cDev{0, 0x22, 0, 0x70, 0x20, 0, 0, 0}
	//port 17-32 reset signals
	qsfpVdevIo[7] = qsfpI2cDev{0, 0x23, 0, 0x70, 0x20, 0, 0, 0}

	qsfpVpageByKeyIo = map[string]uint8{
		"port-" + strconv.Itoa(portOffset+1) + ".qsfp.presence":  0,
		"port-" + strconv.Itoa(portOffset+2) + ".qsfp.presence":  0,
		"port-" + strconv.Itoa(portOffset+3) + ".qsfp.presence":  0,
		"port-" + strconv.Itoa(portOffset+4) + ".qsfp.presence":  0,
		"port-" + strconv.Itoa(portOffset+5) + ".qsfp.presence":  0,
		"port-" + strconv.Itoa(portOffset+6) + ".qsfp.presence":  0,
		"port-" + strconv.Itoa(portOffset+7) + ".qsfp.presence":  0,
		"port-" + strconv.Itoa(portOffset+8) + ".qsfp.presence":  0,
		"port-" + strconv.Itoa(portOffset+9) + ".qsfp.presence":  0,
		"port-" + strconv.Itoa(portOffset+10) + ".qsfp.presence": 0,
		"port-" + strconv.Itoa(portOffset+11) + ".qsfp.presence": 0,
		"port-" + strconv.Itoa(portOffset+12) + ".qsfp.presence": 0,
		"port-" + strconv.Itoa(portOffset+13) + ".qsfp.presence": 0,
		"port-" + strconv.Itoa(portOffset+14) + ".qsfp.presence": 0,
		"port-" + strconv.Itoa(portOffset+15) + ".qsfp.presence": 0,
		"port-" + strconv.Itoa(portOffset+16) + ".qsfp.presence": 0,
		"port-" + strconv.Itoa(portOffset+17) + ".qsfp.presence": 1,
		"port-" + strconv.Itoa(portOffset+18) + ".qsfp.presence": 1,
		"port-" + strconv.Itoa(portOffset+19) + ".qsfp.presence": 1,
		"port-" + strconv.Itoa(portOffset+20) + ".qsfp.presence": 1,
		"port-" + strconv.Itoa(portOffset+21) + ".qsfp.presence": 1,
		"port-" + strconv.Itoa(portOffset+22) + ".qsfp.presence": 1,
		"port-" + strconv.Itoa(portOffset+23) + ".qsfp.presence": 1,
		"port-" + strconv.Itoa(portOffset+24) + ".qsfp.presence": 1,
		"port-" + strconv.Itoa(portOffset+25) + ".qsfp.presence": 1,
		"port-" + strconv.Itoa(portOffset+26) + ".qsfp.presence": 1,
		"port-" + strconv.Itoa(portOffset+27) + ".qsfp.presence": 1,
		"port-" + strconv.Itoa(portOffset+28) + ".qsfp.presence": 1,
		"port-" + strconv.Itoa(portOffset+29) + ".qsfp.presence": 1,
		"port-" + strconv.Itoa(portOffset+30) + ".qsfp.presence": 1,
		"port-" + strconv.Itoa(portOffset+31) + ".qsfp.presence": 1,
		"port-" + strconv.Itoa(portOffset+32) + ".qsfp.presence": 1,
	}

	qsfpVdev[0] = qsfpI2cDev{0, 0x50, 0, 0x70, 0x1, 0, 0x71, 0x1}
	qsfpVdev[1] = qsfpI2cDev{0, 0x50, 0, 0x70, 0x1, 0, 0x71, 0x2}
	qsfpVdev[2] = qsfpI2cDev{0, 0x50, 0, 0x70, 0x1, 0, 0x71, 0x4}
	qsfpVdev[3] = qsfpI2cDev{0, 0x50, 0, 0x70, 0x1, 0, 0x71, 0x8}
	qsfpVdev[4] = qsfpI2cDev{0, 0x50, 0, 0x70, 0x1, 0, 0x71, 0x10}
	qsfpVdev[5] = qsfpI2cDev{0, 0x50, 0, 0x70, 0x1, 0, 0x71, 0x20}
	qsfpVdev[6] = qsfpI2cDev{0, 0x50, 0, 0x70, 0x1, 0, 0x71, 0x40}
	qsfpVdev[7] = qsfpI2cDev{0, 0x50, 0, 0x70, 0x1, 0, 0x71, 0x80}
	qsfpVdev[8] = qsfpI2cDev{0, 0x50, 0, 0x70, 0x2, 0, 0x71, 0x1}
	qsfpVdev[9] = qsfpI2cDev{0, 0x50, 0, 0x70, 0x2, 0, 0x71, 0x2}
	qsfpVdev[10] = qsfpI2cDev{0, 0x50, 0, 0x70, 0x2, 0, 0x71, 0x4}
	qsfpVdev[11] = qsfpI2cDev{0, 0x50, 0, 0x70, 0x2, 0, 0x71, 0x8}
	qsfpVdev[12] = qsfpI2cDev{0, 0x50, 0, 0x70, 0x2, 0, 0x71, 0x10}
	qsfpVdev[13] = qsfpI2cDev{0, 0x50, 0, 0x70, 0x2, 0, 0x71, 0x20}
	qsfpVdev[14] = qsfpI2cDev{0, 0x50, 0, 0x70, 0x2, 0, 0x71, 0x40}
	qsfpVdev[15] = qsfpI2cDev{0, 0x50, 0, 0x70, 0x2, 0, 0x71, 0x80}
	qsfpVdev[16] = qsfpI2cDev{0, 0x50, 0, 0x70, 0x4, 0, 0x71, 0x1}
	qsfpVdev[17] = qsfpI2cDev{0, 0x50, 0, 0x70, 0x4, 0, 0x71, 0x2}
	qsfpVdev[18] = qsfpI2cDev{0, 0x50, 0, 0x70, 0x4, 0, 0x71, 0x4}
	qsfpVdev[19] = qsfpI2cDev{0, 0x50, 0, 0x70, 0x4, 0, 0x71, 0x8}
	qsfpVdev[20] = qsfpI2cDev{0, 0x50, 0, 0x70, 0x4, 0, 0x71, 0x10}
	qsfpVdev[21] = qsfpI2cDev{0, 0x50, 0, 0x70, 0x4, 0, 0x71, 0x20}
	qsfpVdev[22] = qsfpI2cDev{0, 0x50, 0, 0x70, 0x4, 0, 0x71, 0x40}
	qsfpVdev[23] = qsfpI2cDev{0, 0x50, 0, 0x70, 0x4, 0, 0x71, 0x80}
	qsfpVdev[24] = qsfpI2cDev{0, 0x50, 0, 0x70, 0x8, 0, 0x71, 0x1}
	qsfpVdev[25] = qsfpI2cDev{0, 0x50, 0, 0x70, 0x8, 0, 0x71, 0x2}
	qsfpVdev[26] = qsfpI2cDev{0, 0x50, 0, 0x70, 0x8, 0, 0x71, 0x4}
	qsfpVdev[27] = qsfpI2cDev{0, 0x50, 0, 0x70, 0x8, 0, 0x71, 0x8}
	qsfpVdev[28] = qsfpI2cDev{0, 0x50, 0, 0x70, 0x8, 0, 0x71, 0x10}
	qsfpVdev[29] = qsfpI2cDev{0, 0x50, 0, 0x70, 0x8, 0, 0x71, 0x20}
	qsfpVdev[30] = qsfpI2cDev{0, 0x50, 0, 0x70, 0x8, 0, 0x71, 0x40}
	qsfpVdev[31] = qsfpI2cDev{0, 0x50, 0, 0x70, 0x8, 0, 0x71, 0x80}
}
