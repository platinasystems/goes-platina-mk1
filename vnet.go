// Copyright © 2015-2018 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/platinasystems/atsock"
	"github.com/platinasystems/goes/cmd"
	"github.com/platinasystems/goes/lang"
)

type vnetCommand struct{}

func (vnetCommand) String() string { return "vnet" }

func (vnetCommand) Usage() string { return "vnet COMMAND [ARG]..." }

func (vnetCommand) Apropos() lang.Alt {
	return lang.Alt{
		lang.EnUS: "send commands to hidden cli",
	}
}

func (vnetCommand) Close() error { return vnetClose() }

func (vnetCommand) Kind() cmd.Kind { return cmd.DontFork }

func (vnetCommand) Main(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("no COMMAND")
	}
	if err := vnetConnect(); err != nil {
		return err
	}
	return vnetExec(os.Stdout, args...)
}

func (vnetCommand) Help(...string) string {
	err := vnetConnect()
	if err != nil {
		return err.Error()
	}
	buf := new(bytes.Buffer)
	err = vnetExec(buf, "help")
	if err != nil {
		return err.Error()
	} else {
		return buf.String()
	}
}

var vnetConn struct {
	net.Conn
	sync.Mutex
}

func vnetConnect() (err error) {
	vnetConn.Lock()
	defer vnetConn.Unlock()
	if vnetConn.Conn == nil {
		vnetConn.Conn, err = atsock.Dial("vnet")
	}
	return
}

func vnetClose() (err error) {
	vnetConn.Lock()
	defer vnetConn.Unlock()
	if vnetConn.Conn != nil {
		err = vnetConn.Close()
		vnetConn.Conn = nil
	}
	return
}

// Exec runs a vnet cli command and copies output to given io.Writer.
func vnetExec(w io.Writer, args ...string) (err error) {
	var werr error

	// Send cli command to vnet.
	fmt.Fprintf(vnetConn, "%s\n", strings.Join(args, " "))

	// Ignore pipe error e.g. vnet command | head
	signal.Notify(make(chan os.Signal, 1), syscall.SIGPIPE)

	for {
		// First read 32 bit network byte order length.
		var tmp [4]byte
		if _, err = vnetConn.Read(tmp[:]); err != nil {
			return
		}
		if l := int64(binary.BigEndian.Uint32(tmp[:])); l == 0 {
			// Zero length means end of vnet command output.
			break
		} else if werr == nil {
			// Otherwise copy input to output up to first error
			_, werr = io.CopyN(w, vnetConn, l)
		}
	}
	return
}
