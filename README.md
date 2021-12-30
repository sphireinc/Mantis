[![Build and Test](https://github.com/sphireinc/Mantis/actions/workflows/build-and-test.yml/badge.svg?branch=master)](https://github.com/sphireinc/Mantis/actions/workflows/build-and-test.yml)
[![](https://img.shields.io/github/go-mod/go-version/sphireinc/mantis)]()
[![Release Date](https://img.shields.io/github/release-date/sphireinc/mantis)](https://github.com/sphireinc/Mantis/releases/latest)
[![Latest Release](https://img.shields.io/github/v/release/sphireinc/mantis)](https://github.com/sphireinc/Mantis/releases/latest)


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

<p>
  <a href="#sphire-mantis">About</a> •
  <a href="#importing">Importing</a> •
  <a href="#testing">Testing</a>

[//]: # (  <a href="#api-reference">API Reference</a>)
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

[//]: # (# API Reference)

[//]: # ()
[//]: # (* byte)

[//]: # (  * `type ParseErr struct`)

[//]: # (    * `&#40;*ParseErr&#41; String&#40;&#41;`)

[//]: # (    * `&#40;*ParseErr&#41; Error&#40;&#41;`)

[//]: # (  * `type BytesUnit int64`)

[//]: # (    * `&#40;u BytesUnit&#41; Convert&#40;b Bytes&#41;`)

[//]: # (    * `&#40;u BytesUnit&#41; String&#40;&#41;`)

[//]: # (  * `type Bytes int64`)

[//]: # (    * `&#40;b Bytes&#41; Format&#40;s fmt.State, verb rune&#41;`)

[//]: # (    * `&#40;b Bytes&#41; String&#40;&#41;`)

[//]: # (  * `ParseBytes&#40;s string&#41;`)

[//]: # (* cache)

[//]: # (  * `type MemCache struct`)

[//]: # (    * `&#40;m *MemCache&#41; String&#40;&#41;`)

[//]: # (    * `&#40;m *MemCache&#41; Init&#40;&#41;`)

[//]: # (  * `NewMemCache&#40;algorithm memory.Algorithm, capacity int, refreshKey string, cacheTime time.Duration&#41;`)

[//]: # (  * `type BigCache struct`)

[//]: # (    * `&#40;b *BigCache&#41; Init&#40;&#41;`)

[//]: # (* data)

[//]: # (  * `IsTrue&#40;str string&#41;`)

[//]: # (  * `JsonQuery&#40;jsonObj string, query string&#41;`)

[//]: # (  * `DirectoryExists&#40;path string&#41;`)

[//]: # (  * `MapStringStringContains&#40;item map[string]string, key string&#41;`)

[//]: # (* database)

[//]: # (  * `type MySQL struct`)

[//]: # (    * `&#40;m *MySQL&#41; ConfigString&#40;&#41;`)

[//]: # (    * `&#40;m *MySQL&#41; String&#40;&#41;`)

[//]: # (    * `&#40;m *MySQL&#41; Connect&#40;&#41;`)

[//]: # (    * `&#40;m *MySQL&#41; SelectOne&#40;query string, args ...interface{}&#41;`)

[//]: # (    * `&#40;m *MySQL&#41; Select&#40;query string, args ...interface{}&#41;`)

[//]: # (    * `&#40;m *MySQL&#41; Insert&#40;query string, args ...interface{}&#41;`)

[//]: # (    * `&#40;m *MySQL&#41; Update&#40;query string, args ...interface{}&#41;`)

[//]: # (    * `&#40;m *MySQL&#41; Delete&#40;query string, args ...interface{}&#41;`)

[//]: # (  * `type Redis struct`)

[//]: # (    * `&#40;r *Redis&#41; String&#40;&#41;`)

[//]: # (    * `&#40;r *Redis&#41; Init&#40;&#41;`)

[//]: # (    * `&#40;r *Redis&#41; CheckIfConnected&#40;&#41;`)

[//]: # (    * `&#40;r *Redis&#41; Get&#40;key string&#41;`)

[//]: # (    * `&#40;r *Redis&#41; Set&#40;key string, value string, expiration time.Duration&#41;`)

[//]: # (    * `&#40;r *Redis&#41; GetRawConnectionAndContext&#40;&#41;`)

[//]: # (  * `type Neo4j struct`)

[//]: # (    * `&#40;n *Neo4j&#41; String&#40;&#41;`)

[//]: # (    * `&#40;n *Neo4j&#41; Connect&#40;&#41;`)

[//]: # (    * `&#40;n *Neo4j&#41; NewNode&#40;node neoism.Props&#41;`)

[//]: # (    * `&#40;n *Neo4j&#41; CypherQuery&#40;query CypherQuery&#41;`)

[//]: # (    * `&#40;n *Neo4j&#41; TransactCypherQuery&#40;queries []CypherQuery&#41;`)

[//]: # (  * `type CypherQuery struct`)

[//]: # (    * `&#40;c *CypherQuery&#41; String&#40;&#41;`)

[//]: # (* date)

[//]: # (  * `type Date struct`)

[//]: # (    * `&#40;d *Date&#41; String&#40;&#41;`)

[//]: # (    * `&#40;d *Date&#41; DateToString&#40;&#41;`)

[//]: # (  * `CurrentTime&#40;&#41;`)

[//]: # (  * `StringToDate&#40;date string&#41;`)

[//]: # (* encoding)

[//]: # (  * `Base64EncodeStd&#40;data string&#41;`)

[//]: # (  * `Base64EncodeUrl&#40;data string&#41;`)

[//]: # (  * `Base64Decode&#40;encodedData string&#41;`)

[//]: # (* encryption)

[//]: # (  * `type Hash struct`)

[//]: # (    * `&#40;h *Hash&#41; Hash&#40;&#41;`)

[//]: # (* helper)

[//]: # (  * `Reverse&#40;s string&#41;`)

[//]: # (  * `StrConvParseBoolHideError&#40;boolean string&#41;`)

[//]: # (  * `StrConvAtoiWithDefault&#40;intAsString string, defaultValue int&#41;`)

[//]: # (  * `StrConvAtoiWithDefaultTimeDuration&#40;intAsString string, defaultValue int&#41;`)

[//]: # (  * `StringWithDefault&#40;givenValue string, defaultValue string&#41;`)

[//]: # (  * `IntWithDefault&#40;givenValue int, defaultValue int&#41;`)

[//]: # (* http)

[//]: # (  * `type ResponseJsonError struct`)

[//]: # (    * `&#40;r *ResponseJsonError&#41; String&#40;&#41;`)

[//]: # (  * `type ResponseJsonOk struct`)

[//]: # (    * `&#40;r *ResponseJsonOk&#41; String&#40;&#41;`)

[//]: # (  * `type ResponseCodes struct`)

[//]: # (    * `&#40;r *ResponseCodes&#41; String&#40;&#41;`)

[//]: # (  * `GetHTTPResponseCode&#40;code int16&#41;`)

[//]: # (  * `type Request struct`)

[//]: # (    * `&#40;r *Request&#41; String&#40;&#41;`)

[//]: # (    * `&#40;r *Request&#41; Get&#40;&#41;`)

[//]: # (    * `&#40;r *Request&#41; Post&#40;&#41;`)

[//]: # (  * `type Response struct`)

[//]: # (    * `func &#40;r *Response&#41; String&#40;&#41;`)

[//]: # (  * `ParseBodyIntoStruct&#40;r *http.Request, obj interface{}&#41;`)

[//]: # (  * `GetBody&#40;r *http.Request&#41;`)

[//]: # (  * `GetQueryParameter&#40;r *http.Request, key string&#41;`)

[//]: # (  * `GetQueryParameters&#40;r *http.Request&#41;`)

[//]: # (  * `ParseUrl&#40;rawurl string&#41;`)

[//]: # (* log)

[//]: # (  * `type Log struct`)

[//]: # (    * `&#40;l *Log&#41; String&#40;&#41;`)

[//]: # (    * `&#40;l *Log&#41; Write&#40;msg string&#41;`)

[//]: # (    * `&#40;l *Log&#41; LogHTTPRequest&#40;name string, w http.ResponseWriter, r *http.Request&#41;`)

[//]: # (    * `&#40;l *Log&#41; HandleError&#40;message string, err error&#41;`)

[//]: # (    * `&#40;l *Log&#41; HandleFatalError&#40;err error&#41;`)

[//]: # (    * `&#40;l *Log&#41; JSONMarshalAndLogError&#40;message string, err error&#41;`)

[//]: # (  * `New&#40;filename string&#41;`)

[//]: # (  * `JSONMarshalError&#40;err error&#41;`)

[//]: # (* uuid)

[//]: # (  * `type UUID [16]byte`)

[//]: # (    * `&#40;u UUID&#41; Version&#40;&#41;`)

[//]: # (    * `&#40;u *UUID&#41; SetVersion&#40;ver byte&#41;`)

[//]: # (    * `&#40;u *UUID&#41; SetDCESecurity&#40;domain byte, id uint32&#41;`)

[//]: # (    * `&#40;u *UUID&#41; DCESecurity&#40;&#41;`)

[//]: # (  * `type Variant byte`)

[//]: # (    * `&#40;u UUID&#41; Variant&#40;&#41;`)

[//]: # (    * `&#40;u *UUID&#41; SetVariant&#40;v Variant&#41;`)

[//]: # (    * `&#40;u UUID&#41; Time&#40;&#41;`)

[//]: # (    * `&#40;u UUID&#41; String&#40;&#41;`)

[//]: # (    * `&#40;u UUID&#41; Format&#40;s fmt.State, verb rune&#41;`)

[//]: # (    * `&#40;u *UUID&#41; UnmarshalText&#40;text []byte&#41;`)

[//]: # (    * `&#40;u *UUID&#41; MarshalText&#40;&#41;`)

[//]: # (    * `&#40;u *UUID&#41; Equals&#40;o *UUID&#41;`)

[//]: # (  * `GenerateV1&#40;&#41;`)

[//]: # (  * `GenerateV2&#40;domain byte, id uint32&#41;`)

[//]: # (  * `GenerateV3&#40;ns UUID, n []byte&#41;`)

[//]: # (  * `GenerateV4&#40;&#41;`)

[//]: # (  * `GenerateV5&#40;ns UUID, n []byte&#41;`)

[//]: # (  * `GenerateV4String&#40;&#41;`)

[//]: # (  * `MustParseUUIDString&#40;s string&#41;`)

[//]: # (  * `ParseUUIDString&#40;s string&#41;`)
