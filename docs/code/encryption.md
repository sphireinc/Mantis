<!-- Code generated by gomarkdoc. DO NOT EDIT -->

[![Go Report Card](https://goreportcard.com/badge/github.com/sphireinc/mantis)](https://goreportcard.com/report/github.com/sphireinc/mantis)
[![Go Reference](https://pkg.go.dev/badge/github.com/sphireinc/mantis.svg)](https://pkg.go.dev/github.com/sphireinc/mantis)



# encryption

```go
import "github.com/sphireinc/mantis/encryption"
```

## Index

- [Constants](<#constants>)
- [func CreateRandomBytes\(bytes int\) \[\]byte](<#CreateRandomBytes>)
- [func CreateRandomString\(bytes int\) string](<#CreateRandomString>)
- [type MHash](<#MHash>)
  - [func New\(input string, algorithm int8\) \*MHash](<#New>)
  - [func \(h \*MHash\) Algorithm\(\) \(int8, string\)](<#MHash.Algorithm>)
  - [func \(h \*MHash\) GetInput\(\) string](<#MHash.GetInput>)
  - [func \(h \*MHash\) GetOutput\(\) string](<#MHash.GetOutput>)
  - [func \(h \*MHash\) Hash\(\)](<#MHash.Hash>)
  - [func \(h \*MHash\) IsHashed\(\) bool](<#MHash.IsHashed>)
  - [func \(h \*MHash\) MarshalJSON\(\) \(\[\]byte, error\)](<#MHash.MarshalJSON>)


## Constants

<a name="Md5"></a>Md5 are our enumerators Sha224 Sha256 Sha384 Sha512 Sha512224 Sha512256 Hmac512

```go
const (
    Md5 int8 = iota
    Sha224
    Sha256
    Sha384
    Sha512
    Sha512224
    Sha512256
    Hmac512
)
```

<a name="CreateRandomBytes"></a>
## func [CreateRandomBytes](<https://github.com/sphireinc/mantis/blob/master/encryption/encryption.go#L140>)

```go
func CreateRandomBytes(bytes int) []byte
```

CreateRandomBytes creates a random bytes of bytes int

<a name="CreateRandomString"></a>
## func [CreateRandomString](<https://github.com/sphireinc/mantis/blob/master/encryption/encryption.go#L135>)

```go
func CreateRandomString(bytes int) string
```

CreateRandomString generates a random string of n bytes

<a name="MHash"></a>
## type [MHash](<https://github.com/sphireinc/mantis/blob/master/encryption/encryption.go#L37-L42>)

MHash our input/output tracking struct

```go
type MHash struct {
    input     string
    isHashed  bool
    Output    string
    algorithm int8
}
```

<a name="New"></a>
### func [New](<https://github.com/sphireinc/mantis/blob/master/encryption/encryption.go#L54>)

```go
func New(input string, algorithm int8) *MHash
```

New returns an instance of MHash given our input and algorithm

<a name="MHash.Algorithm"></a>
### func \(\*MHash\) [Algorithm](<https://github.com/sphireinc/mantis/blob/master/encryption/encryption.go#L67>)

```go
func (h *MHash) Algorithm() (int8, string)
```

Algorithm returns the chosen algorithm as a string and int

<a name="MHash.GetInput"></a>
### func \(\*MHash\) [GetInput](<https://github.com/sphireinc/mantis/blob/master/encryption/encryption.go#L90>)

```go
func (h *MHash) GetInput() string
```

GetInput returns the initial input

<a name="MHash.GetOutput"></a>
### func \(\*MHash\) [GetOutput](<https://github.com/sphireinc/mantis/blob/master/encryption/encryption.go#L95>)

```go
func (h *MHash) GetOutput() string
```

GetOutput returns the hashed output

<a name="MHash.Hash"></a>
### func \(\*MHash\) [Hash](<https://github.com/sphireinc/mantis/blob/master/encryption/encryption.go#L100>)

```go
func (h *MHash) Hash()
```

Hash performs our hash, fills in Output, and unsets input

<a name="MHash.IsHashed"></a>
### func \(\*MHash\) [IsHashed](<https://github.com/sphireinc/mantis/blob/master/encryption/encryption.go#L62>)

```go
func (h *MHash) IsHashed() bool
```

IsHashed tells us whether our MHash has been MHash\(\)'d

<a name="MHash.MarshalJSON"></a>
### func \(\*MHash\) [MarshalJSON](<https://github.com/sphireinc/mantis/blob/master/encryption/encryption.go#L45>)

```go
func (h *MHash) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the JSON encoding interface

Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)