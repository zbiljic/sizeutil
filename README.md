# sizeutil

[![GoDoc](https://godoc.org/github.com/zbiljic/sizeutil?status.svg)](https://godoc.org/github.com/zbiljic/sizeutil)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/zbiljic/sizeutil/master/LICENSE)

This package contains functions for working with byte units in Go (golang):

*   Parses human readable string to byte unit
*   Converts from one byte unit to another
*   Formats byte unit to human readable format

All internal numberic types are `int64`.

## Installation

```bash
go get github.com/zbiljic/sizeutil
```

## Example:

```go
sizeutil.ParseSize("256MB")             // returns &sizeutil.Size{count: 256, unit: 8388608}
sizeutil.Bytes(268435456).ToMegabytes() // returns 256
sizeutil.Megabytes(256).String()        // returns "256 megabytes"
```

See the [reference][] for more info.

[reference]: http://godoc.org/github.com/zbiljic/sizeutil

---

Copyright © 2017 Nemanja Zbiljić
