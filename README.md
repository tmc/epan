epan
============

cgo bindings for libwireshark's epan packet dissection library

* status: incomplete -- contributions welcome!
* license: isc

Installation
------------

Archlinux:

```sh
  $ pacman -S wireshark-cli
  $ go get github.com/tmc/epan
```

Debian/Ubuntu:

```sh
  $ apt-get install libwireshark-dev
  $ go get github.com/tmc/epan
```
MacOS:

```sh
  $ brew install wireshark --with-headers
  $ go get github.com/tmc/epan
```

to list recognized protocols:
```sh
  $ go install github.com/tmc/epan/examples/list_protocols
  $ list_protocols
```
