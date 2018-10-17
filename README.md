This is a GO Embedded System for Platina Systems' *mark 1* packet switches.

To build this source you'll first need to extract the driver plugin from and
existing `goes-platina-mk1` binary.

```console
$ unzip goes-platina-mk1 fe1.so
```

You may then rebuild the binary like this,

```console
/$ go build -ldflags "-X 'main.Version=$(git describe)'"
$ cat fe1.zip >> goes-platina-mk1
$ zip -q -A goes-platina-mk1
```

Then copy and install this on a *mark 1* switch,

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
