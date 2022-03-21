[![Release Version](https://img.shields.io/github/v/release/sphireinc/mantis)](https://github.com/sphireinc/Mantis/releases/latest)
[![Build Status](https://github.com/sphireinc/Mantis/actions/workflows/build.yml/badge.svg?branch=master)](https://github.com/sphireinc/Mantis/actions/workflows/build-and-test.yml)
[![Go Version](https://img.shields.io/github/go-mod/go-version/sphireinc/mantis)]()
[![Release Date](https://img.shields.io/github/release-date/sphireinc/mantis)](https://github.com/sphireinc/Mantis/releases/latest)
[![Codacy Security Scan](https://github.com/sphireinc/Mantis/actions/workflows/codacy-analysis.yml/badge.svg?branch=master)](https://github.com/sphireinc/Mantis/actions/workflows/codacy-analysis.yml)
[![CodeQL](https://github.com/sphireinc/Mantis/actions/workflows/codeql-analysis.yml/badge.svg?branch=master)](https://github.com/sphireinc/Mantis/actions/workflows/codeql-analysis.yml)

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

Mantis is pre-production - use at your own risk. If you do decide to use it, be prepared for breaking changes in the API.

# Sphire Mantis

Mantis is a common helper library within the Sphire ecosystem. Packages are largely developed 
to have no effect on the parent application. Helper functions are (or will be in some cases) 
developed to accept, act on, and return data with no side effects.

# Usage

Running `go get github.com/sphireinc/mantis` will fetch the mantis project within your Go project.

Imports work at a package level. If you'd like to use the date package, please do:

```go
package main 

import (
	mantisDate `github.com/sphireinc/mantis/date`
)

func main(){ 
	var d mantisDate.Date = mantisDate.CurrentTime()
	fmt.Println(d.String())
}
```

# Tests

Each package can be tested independently via `go test`, or all packages can be tested from 
the root via `go test ./...`


# Contributing

Please find our contribution guidelines within [CONTRIBUTING.md](https://github.com/sphireinc/Mantis/blob/master/CONTRIBUTING.md)

# Thanks

Thanks to all below for their contributions, inspiration, or otherwise:

* https://github.com/dec0dOS/amazing-github-template for the `.github ISSUE_TEMPLATE`
* https://github.com/junosuarez/CONTRIBUTING.md for the `CONTRIBUTING.md` template