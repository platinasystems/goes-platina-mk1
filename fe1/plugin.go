// Copyright © 2015-2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

// This is Platina's fe1 dynamic library; build it with,
//	go build -ldflags "-X 'main.Version=$(git describe)'" -buildmode=plugin
//	zip ../fe1.zip fe1.so
// Test with,
//	go test -c
//	sudo GOPATH=$GOPATH ./fe1.test [-test.help]
package main

import (
	"github.com/platinasystems/fe1"
	firmware "github.com/platinasystems/firmware-fe1a"
	"github.com/platinasystems/go/vnet"
	platform "github.com/platinasystems/go/vnet/platforms/fe1"
)

const (
	ldflags = `-ldflags "-X 'main.Version=$(git describe)'"`
	version = "FIXME with, go build " + ldflags + " -buildmode=plugin"
)

var Version = version

func Licenses() map[string]string {
	return map[string]string{
		"fe1":      fe1.License,
		"firmware": firmware.License,
	}
}
func Patents() map[string]string {
	return map[string]string{
		"fe1": fe1.Patents,
	}
}

func Versions() map[string]string {
	return map[string]string{
		"fe1": Version,
	}
}

func Init(v *vnet.Vnet, p *platform.Platform) {
	fe1.Init(v, p)
}

func AddPlatform(v *vnet.Vnet, p *platform.Platform) {
	fe1.AddPlatform(v, p)
}
