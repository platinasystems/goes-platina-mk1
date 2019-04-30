This is a GO Embedded System for Platina Systems' *mark 1* packet switches.

To build this from source you'll first need to extract the driver from an
existing program binary.

```console
$ unzip goes-platina-mk1 vnet-platina-mk1
$ zip drivers vnet-platina-mk1
```

Or with NDA access to the driver source, build it with,

```console
$ go get github.com/platinasystems/vnet-platina-mk1@VERSION
$ zip -j drivers $GOPATH/bin/vnet-platina-mk1
```

Then build the goes program and append the zipped plugin.

```console
$ go get github.com/platinasystems/goes-platina-mk1@VERSION
$ cat $GOPATH/bin/goes-platina-mk1 drivers.zip >> goes-platina-mk1
$ zip -q -A goes-platina-mk1
```

Install this on a *mark 1* switch with,

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

---

*&copy; 2015-2018 Platina Systems, Inc. All rights reserved.
Use of this source code is governed by this BSD-style [LICENSE].*

[LICENSE]: LICENSE
[errata]: docs/Errata.md
