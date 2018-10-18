This is the plugin driver for Platina Systems' *mark 1* packet switches.

With NDA access to the imported source, build with,

```console
$ go build -ldflags "-X 'main.Version=$(git describe)'" -buildmode=plugin
$ zip ../fe1.zip fe1.so
```

---

*&copy; 2015-2018 Platina Systems, Inc. All rights reserved.
Use of this source code is governed by this BSD-style [LICENSE].*

[LICENSE]: ../LICENSE
