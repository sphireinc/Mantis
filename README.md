[![Release Version](https://img.shields.io/github/v/release/sphireinc/mantis)](https://github.com/sphireinc/Mantis/releases/latest)
[![Release Date](https://img.shields.io/github/release-date/sphireinc/mantis)](https://github.com/sphireinc/Mantis/releases/latest)
[![Build Status](https://github.com/sphireinc/Mantis/actions/workflows/build.yml/badge.svg?branch=master)](https://github.com/sphireinc/Mantis/actions/workflows/build-and-test.yml)
[![Go Version](https://img.shields.io/github/go-mod/go-version/sphireinc/mantis)](https://github.com/sphireinc/Mantis/releases/latest)
[![License](https://img.shields.io/github/license/sphireinc/mantis)](https://github.com/sphireinc/Mantis/releases/latest)
[![Codacy and CodeQL](https://github.com/sphireinc/Mantis/actions/workflows/analyze.yml/badge.svg?branch=master)](https://github.com/sphireinc/Mantis/actions/workflows/codeql-analysis.yml)

<p>
    <img src="https://raw.githubusercontent.com/sphireinc/Mantis/master/_logo/mantis_logo.png" alt="Sphire Mantis Logo"/>
</p>

> A broadly featured Go helper library with standalone packages

<p>
  <a target="_blank" href="https://github.com/sphireinc/Mantis/tree/master/byte"><img src="https://img.shields.io/badge/Byte-brightgreen" alt=""/></a>&nbsp;
  <a target="_blank" href="https://github.com/sphireinc/Mantis/tree/master/cache"><img src="https://img.shields.io/badge/Cache-blue" alt=""/></a>&nbsp;
  <a target="_blank" href="https://github.com/sphireinc/Mantis/tree/master/data"><img src="https://img.shields.io/badge/Data-orangered" alt=""/></a>&nbsp;
  <a target="_blank" href="https://github.com/sphireinc/Mantis/tree/master/database"><img src="https://img.shields.io/badge/Database-violet" alt=""/></a>&nbsp;
  <a target="_blank" href="https://github.com/sphireinc/Mantis/tree/master/date"><img src="https://img.shields.io/badge/Date-informational" alt=""/></a>&nbsp;
  <a target="_blank" href="https://github.com/sphireinc/Mantis/tree/master/encoding"><img src="https://img.shields.io/badge/Encoding-brightgreen" alt=""/></a>&nbsp;
  <a target="_blank" href="https://github.com/sphireinc/Mantis/tree/master/encryption"><img src="https://img.shields.io/badge/Encryption-orangered" alt=""/></a>&nbsp;
  <a target="_blank" href="https://github.com/sphireinc/Mantis/tree/master/helper"><img src="https://img.shields.io/badge/Helper-important" alt=""/></a>&nbsp;
  <a target="_blank" href="https://github.com/sphireinc/Mantis/tree/master/http"><img src="https://img.shields.io/badge/HTTP-critical" alt=""/></a>&nbsp;
  <a target="_blank" href="https://github.com/sphireinc/Mantis/tree/master/log"><img src="https://img.shields.io/badge/Log-blue" alt=""/></a>&nbsp;
  <a target="_blank" href="https://github.com/sphireinc/Mantis/tree/master/uuid"><img src="https://img.shields.io/badge/UUID-lightgrey" alt=""/></a>&nbsp;
</p>


# Warning

Mantis is pre-production - use at your own risk.

# Mantis

Mantis is a collection of helper libraries written in order to reduce
the need to code common patterns and functionality. Subpackages are developed, when possible, 
to have zero side effects on the parent application - helper functions are (or in some cases will be)
written in such a manner that they accept, act on, and return data.

Mantis, much like its namesake ([Jet's Go-Mantis](https://github.com/jet/go-mantis) library), is the
"standard library" for Sphire's Golang codebase, used heavily in projects like Sphire Core and Sphire Codex. 

# Usage

Running `go get github.com/sphireinc/mantis` will fetch the mantis project within your Go project.

Imports work at a package level. If you'd like to use the date package, please do:

```go
package main 

import (
	mantisDate `github.com/sphireinc/mantis/date`
)

func main(){ 
	datem := mantisDate.CurrentTime() // return type: mantisDate.Date
	fmt.Println(datem.String())
}
```

```bash
$ go run main.go
{
  "year": 2009,
  "month": 11,
  "day": 10,
  "hour": 23,
  "minute": 0,
  "second": 0,
  "nanosecond": 0,
  "unix": 1257894000,
  "week_day": 2,
  "year_day": 314
}
```

# Tests

Each package can be tested independently via `go test`, or all packages can be tested from 
the root via `go test ./...`

# Local Development

You should ideally install these packages:

```bash
go install golang.org/x/lint/golint@latest
go install github.com/securego/gosec/v2/cmd/gosec@latest
go install golang.org/x/tools/cmd/goimports@latest
go install honnef.co/go/tools/cmd/staticcheck@latest
```

Then you should run these commands, ideally as a pre-commit check:

```bash
go fmt
go vet
golint package_name
staticcheck
goimports -v -e -w package_name
```

# Contributing

Please find our contribution guidelines within [CONTRIBUTING.md](https://github.com/sphireinc/Mantis/blob/master/CONTRIBUTING.md)

# Thanks

Thanks to all below for their contributions, inspiration, or otherwise:

* https://github.com/dec0dOS/amazing-github-template for the `.github ISSUE_TEMPLATE`
* https://github.com/junosuarez/CONTRIBUTING.md for the `CONTRIBUTING.md` template