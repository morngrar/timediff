# 🌳 timediff - a tool for calculating clock time differences

[![GoDoc](https://godoc.org/github.com/morngrar/timediff?status.svg)](https://godoc.org/github.com/morngrar/timediff)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)

## Install

This command can be installed as a standalone program or composed into a
Bonzai command tree.

Standalone

```
go install github.com/morngrar/timediff/cmd/timediff@latest
```

Composed

```go
package z

import (
	Z "github.com/rwxrob/bonzai/z"
	timediff "github.com/morngrar/timediff"
)

var Cmd = &Z.Cmd{
	Name:     `z`,
	Commands: []*Z.Cmd{help.Cmd, timediff.Cmd},
}
```

## Tab Completion

To activate bash completion just use the `complete -C` option from your
`.bashrc` or command line. There is no messy sourcing required. All the
completion is done by the program itself.

```
complete -C timediff timediff
```

If you don't have bash or tab completion check use the shortcut
commands instead.

## Embedded Documentation

All documentation (like manual pages) has been embedded into the source
code of the application. See the source or run the program with help to
access it.


## Other Examples

* <https://github.com/rwxrob/z> - the one that started it all
