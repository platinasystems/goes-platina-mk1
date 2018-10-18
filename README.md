This is a GO Embedded System for Platina Systems' *mark 1* packet switches.

To build this source you'll first need to extract the driver plugin from an
existing program binary.

```console
$ unzip goes-platina-mk1 fe1.so
$ zip fe1.zip fe1.so
```

Or with NDA access to the plugin source, build it with,

```console
$ cd fe1
$ go build -ldflags "-X 'main.Version=$(git describe)'" -buildmode=plugin
$ zip ../fe1.zip fe1.so
$ cd ..
```

Then build the program and append the zipped plugin.

```console
$ go build -ldflags "-X 'main.Version=$(git describe)'"
$ cat fe1.zip >> goes-platina-mk1
$ zip -q -A goes-platina-mk1
```

Install this on a *mark 1* switchi with,

```console
$ sudo ./goes-platina-mk1 install
```

To enable BASH completion after install,

```console
. /usr/share/bash-completion/completions/goes
```

To see the commands available on the installed MACHINE,

```console
$ goes help
```

Or,

```console
$ goes
goes> help
```

Then `man` any of the listed commands or `man cli` to see how to use the
command line interface.

See also [errata].

With NDA access to the imported plugin source, you may test with,

```console
$ go test -c
$ sudo GOPATH=$GOPATH ./goes-platina-mk1.test [-test.help]
```

To run unit tests, loopback 6 pairs for ports and edit the configuration
as follows:
```console
$ editor testdata/netport.yaml
$ git update-index --assume-unchanged testdata/netport.yaml
$ editor testdata/ethtool.yaml
$ git update-index --assume-unchanged testdata/ethtool.yaml
```

Test Options:

```console
-test.alpha	this is a zero based alpha system
-test.dryrun	don't run, just print test names
-test.main	internal flag to run given goes command
-test.pause	enable progromatic pause to start debugger
-test.run=Test/PATTERN
		run the matching tests
-test.v		verbose
-test.vv	log test.Program output
-test.vvv	log test.Program execution
```

For example:
```console
sudo GOPATH=$GOPATH ./goes-platina-mk1.test -test.run Test/docker/frr/ospf/eth
sudo GOPATH=$GOPATH ./goes-platina-mk1.test -test.run Test/docker/frr/.*/eth
sudo GOPATH=$GOPATH ./goes-platina-mk1.test -test.run Test/.*/.*/.*/vlan
```

To list all tests:
```console
$ ./goes-platina-mk1.test -test.dryrun
Test
...
PASS
```

---

*&copy; 2015-2018 Platina Systems, Inc. All rights reserved.
Use of this source code is governed by this BSD-style [LICENSE].*

[LICENSE]: LICENSE
[errata]: docs/Errata.md
