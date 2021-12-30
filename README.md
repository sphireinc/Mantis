[![Build and Test](https://github.com/sphireinc/Mantis/actions/workflows/build-and-test.yml/badge.svg?branch=master)](https://github.com/sphireinc/Mantis/actions/workflows/build-and-test.yml)
[![](https://img.shields.io/github/go-mod/go-version/sphireinc/mantis)]()
[![Release Date](https://img.shields.io/github/release-date/sphireinc/mantis)](https://github.com/sphireinc/Mantis/releases/latest)
[![Latest Release](https://img.shields.io/github/v/release/sphireinc/mantis)](https://github.com/sphireinc/Mantis/releases/latest)


<p>
    <img src="https://raw.githubusercontent.com/sphireinc/Mantis/master/_logo/mantis_logo.png" alt="Sphire Mantis Logo"/>
</p>

> A broadly featured Go helper library with standalone packages

<p>
  <a target="_blank" href="https://github.com/sphireinc/Mantis/tree/master/byte"><img src="https://img.shields.io/badge/Bytes-brightgreen" alt=""/></a>&nbsp;
  <a target="_blank" href="https://github.com/sphireinc/Mantis/tree/master/cache"><img src="https://img.shields.io/badge/Cache-blue" alt=""/></a>&nbsp;
  <a target="_blank" href="https://github.com/sphireinc/Mantis/tree/master/data"><img src="https://img.shields.io/badge/Data-orangered" alt=""/></a>&nbsp;
  <a target="_blank" href="https://github.com/sphireinc/Mantis/tree/master/database"><img src="https://img.shields.io/badge/Database-violet" alt=""/></a>&nbsp;
  <a target="_blank" href="https://github.com/sphireinc/Mantis/tree/master/date"><img src="https://img.shields.io/badge/Date-informational" alt=""/></a>&nbsp;
  <a target="_blank" href="https://github.com/sphireinc/Mantis/tree/master/helper"><img src="https://img.shields.io/badge/Helper-important" alt=""/></a>&nbsp;
  <a target="_blank" href="https://github.com/sphireinc/Mantis/tree/master/http"><img src="https://img.shields.io/badge/HTTP-critical" alt=""/></a>&nbsp;
  <a target="_blank" href="https://github.com/sphireinc/Mantis/tree/master/log"><img src="https://img.shields.io/badge/Log-blue" alt=""/></a>&nbsp;
  <a target="_blank" href="https://github.com/sphireinc/Mantis/tree/master/uuid"><img src="https://img.shields.io/badge/UUID-lightgrey" alt=""/></a>&nbsp;
</p>

<p>
  <a href="#sphire-mantis">About</a> •
  <a href="#importing">Importing</a> •
  <a href="#testing">Testing</a> •
  <a href="#api-reference">API Reference</a>
</p>

# Warning

Mantis is pre-production - use at your own risk. If you do decide to use it, be prepared for breaking changes in the API.

# Sphire Mantis

Mantis is a common helper library within the Sphire ecosystem. Packages are largely developed 
to have no effect on the parent application. Helper functions are (or will be in some cases) 
developed to accept, act on, and return data with no side effects.

# Importing

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

# Testing

Each package can be tested independently via `go test`, or all packages can be tested from 
the root via `go test ./...`

# API Reference

* `byte`
  * `type ParseErr struct`
    * `(*ParseErr) String()`
    * `(*ParseErr) Error()`
  * `type BytesUnit int64`
    * `(u BytesUnit) Convert(b Bytes)`
    * `(u BytesUnit) String()`
  * `type Bytes int64`
    * `(b Bytes) Format(s fmt.State, verb rune)`
    * `(b Bytes) String()`
  * `ParseBytes(s string)`
* `cache`
  * `type MemCache struct`
    * `(m *MemCache) String()`
    * `(m *MemCache) Init()`
  * `NewMemCache(algorithm memory.Algorithm, capacity int, refreshKey string, cacheTime time.Duration)`
  * `type BigCache struct`
    * `(b *BigCache) Init()`
* `data`
  * `IsTrue(str string)`
  * `JsonQuery(jsonObj string, query string)`
  * `DirectoryExists(path string)`
  * `MapStringStringContains(item map[string]string, key string)`
* `database`
  * `type MySQL struct`
    * `(m *MySQL) ConfigString()`
    * `(m *MySQL) String()`
    * `(m *MySQL) Connect()`
    * `(m *MySQL) SelectOne(query string, args ...interface{})`
    * `(m *MySQL) Select(query string, args ...interface{})`
    * `(m *MySQL) Insert(query string, args ...interface{})`
    * `(m *MySQL) Update(query string, args ...interface{})`
    * `(m *MySQL) Delete(query string, args ...interface{})`
  * `type Redis struct`
    * `(r *Redis) String()`
    * `(r *Redis) Init()`
    * `(r *Redis) CheckIfConnected()`
    * `(r *Redis) Get(key string)`
    * `(r *Redis) Set(key string, value string, expiration time.Duration)`
    * `(r *Redis) GetRawConnectionAndContext()`
  * `type Neo4j struct`
    * `(n *Neo4j) String()`
    * `(n *Neo4j) Connect()`
    * `(n *Neo4j) NewNode(node neoism.Props)`
    * `(n *Neo4j) CypherQuery(query CypherQuery)`
    * `(n *Neo4j) TransactCypherQuery(queries []CypherQuery)`
  * `type CypherQuery struct`
    * `(c *CypherQuery) String()`
* `date`
  * `type Date struct`
    * `(d *Date) String()`
    * `(d *Date) DateToString()`
  * `CurrentTime()`
  * `StringToDate(date string)`
* `encoding`
  * `Base64EncodeStd(data string)`
  * `Base64EncodeUrl(data string)`
  * `Base64Decode(encodedData string)`
* `encryption`
  * `type Hash struct`
    * `(h *Hash) Hash()`
* `helper`
  * `Reverse(s string)`
  * `StrConvParseBoolHideError(boolean string)`
  * `StrConvAtoiWithDefault(intAsString string, defaultValue int)`
  * `StrConvAtoiWithDefaultTimeDuration(intAsString string, defaultValue int)`
  * `StringWithDefault(givenValue string, defaultValue string)`
  * `IntWithDefault(givenValue int, defaultValue int)`
* `http`
  * `type ResponseJsonError struct`
    * `(r *ResponseJsonError) String()`
  * `type ResponseJsonOk struct`
    * `(r *ResponseJsonOk) String()`
  * `type ResponseCodes struct`
    * `(r *ResponseCodes) String()`
  * `GetHTTPResponseCode(code int16)`
  * `type Request struct`
    * `(r *Request) String()`
    * `(r *Request) Get()`
    * `(r *Request) Post()`
  * `type Response struct`
    * `func (r *Response) String()`
  * `ParseBodyIntoStruct(r *http.Request, obj interface{})`
  * `GetBody(r *http.Request)`
  * `GetQueryParameter(r *http.Request, key string)`
  * `GetQueryParameters(r *http.Request)`
  * `ParseUrl(rawurl string)`
* `log`
  * `type Log struct`
    * `(l *Log) String()`
    * `(l *Log) Write(msg string)`
    * `(l *Log) LogHTTPRequest(name string, w http.ResponseWriter, r *http.Request)`
    * `(l *Log) HandleError(message string, err error)`
    * `(l *Log) HandleFatalError(err error)`
    * `(l *Log) JSONMarshalAndLogError(message string, err error)`
  * `New(filename string)`
  * `JSONMarshalError(err error)`
* `uuid`
  * `type UUID [16]byte`
    * `(u UUID) Version()`
    * `(u *UUID) SetVersion(ver byte)`
    * `(u *UUID) SetDCESecurity(domain byte, id uint32)`
    * `(u *UUID) DCESecurity()`
  * `type Variant byte`
    * `(u UUID) Variant()`
    * `(u *UUID) SetVariant(v Variant)`
    * `(u UUID) Time()`
    * `(u UUID) String()`
    * `(u UUID) Format(s fmt.State, verb rune)`
    * `(u *UUID) UnmarshalText(text []byte)`
    * `(u *UUID) MarshalText()`
    * `(u *UUID) Equals(o *UUID)`
  * `GenerateV1()`
  * `GenerateV2(domain byte, id uint32)`
  * `GenerateV3(ns UUID, n []byte)`
  * `GenerateV4()`
  * `GenerateV5(ns UUID, n []byte)`
  * `GenerateV4String()`
  * `MustParseUUIDString(s string)`
  * `ParseUUIDString(s string)`
