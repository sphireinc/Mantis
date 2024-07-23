<!-- Code generated by gomarkdoc. DO NOT EDIT -->

[![Go Report Card](https://goreportcard.com/badge/github.com/sphireinc/mantis)](https://goreportcard.com/report/github.com/sphireinc/mantis)
[![Go Reference](https://pkg.go.dev/badge/github.com/sphireinc/mantis.svg)](https://pkg.go.dev/github.com/sphireinc/mantis)



# byte

```go
import "github.com/sphireinc/mantis/byte"
```

## Index

- [Constants](<#constants>)
- [Variables](<#variables>)
- [type Bytes](<#Bytes>)
  - [func ParseBytes\(s string\) \(Bytes, error\)](<#ParseBytes>)
  - [func \(b Bytes\) Format\(s fmt.State, verb rune\)](<#Bytes.Format>)
  - [func \(b Bytes\) String\(\) string](<#Bytes.String>)
- [type BytesUnit](<#BytesUnit>)
  - [func \(u BytesUnit\) Convert\(b Bytes\) float64](<#BytesUnit.Convert>)
  - [func \(u BytesUnit\) String\(\) string](<#BytesUnit.String>)
- [type ParseErr](<#ParseErr>)
  - [func \(e \*ParseErr\) Error\(\) string](<#ParseErr.Error>)
  - [func \(e \*ParseErr\) String\(\) string](<#ParseErr.String>)


## Constants

<a name="B"></a>

```go
const (
    // B byte (Unit)
    B   BytesUnit = 1
    // KB Kilobyte (SI) - 1000 bytes
    KB  BytesUnit = 1000
    // MB Megabyte (SI) - 1e6 bytes
    MB  BytesUnit = 1e6
    // GB Gigabyte (SI) - 1e9 bytes
    GB  BytesUnit = 1e9
    // TB Terabyte (SI) - 1e12 bytes
    TB  BytesUnit = 1e12
    // PB Petabyte (SI) - 1e15 bytes
    PB  BytesUnit = 1e15

    // KiB Kibibyte (NIST) - 1024 bytes
    KiB BytesUnit = 1024
    // MiB Mebibyte (NIST) - 2^20 bytes
    MiB = 1024 * KiB
    // GiB Gibibyte (NIST) - 2^30 bytes
    GiB = 1024 * MiB
    // TiB Tebibyte (NIST) - 2^40 bytes
    TiB = 1024 * GiB
    // PiB Pebibyte (NIST) - 2^50 bytes
    PiB = 1024 * TiB
)
```

## Variables

<a name="byteUnitToStr"></a>

```go
var byteUnitToStr = map[BytesUnit]string{
    B:   "B",
    KB:  "KB",
    KiB: "KiB",
    MB:  "MB",
    MiB: "MiB",
    GB:  "GB",
    GiB: "GiB",
    TB:  "TB",
    TiB: "TiB",
    PB:  "PB",
    PiB: "PiB",
}
```

<a name="bytesSIUnitThresholds"></a>

```go
var bytesSIUnitThresholds = []struct {
    unit BytesUnit
    less Bytes
}{
    {B, 5e2},
    {KB, 5e5},
    {MB, 5e8},
    {GB, 5e11},
    {TB, 5e14},
    {PB, 0x7FFFFFFFFFFFFFFF},
}
```

<a name="bytesUnitThresholds"></a>

```go
var bytesUnitThresholds = []struct {
    unit BytesUnit
    less Bytes
}{
    {B, 0x400},
    {KiB, Bytes(uint64(MiB) >> 4)},
    {MiB, Bytes(uint64(GiB) >> 4)},
    {GiB, Bytes(uint64(TiB) >> 4)},
    {TiB, Bytes(uint64(PiB) >> 4)},
    {PiB, 0x7FFFFFFFFFFFFFFF},
}
```

<a name="strToByteUnit"></a>

```go
var strToByteUnit = map[string]BytesUnit{
    "":    B,
    "B":   B,
    "KB":  KB,
    "KIB": KiB,
    "MB":  MB,
    "MIB": MiB,
    "GB":  GB,
    "GIB": GiB,
    "TB":  TB,
    "TIB": TiB,
    "PB":  PB,
    "PIB": PiB,
}
```

<a name="Bytes"></a>
## type [Bytes](<https://github.com/sphireinc/mantis/blob/master/byte/bytes.go#L9>)

Bytes is a formattable byte count

```go
type Bytes int64
```

<a name="ParseBytes"></a>
### func [ParseBytes](<https://github.com/sphireinc/mantis/blob/master/byte/parse.go#L26>)

```go
func ParseBytes(s string) (Bytes, error)
```

ParseBytes parses the given string with byte units into a Byte count

<a name="Bytes.Format"></a>
### func \(Bytes\) [Format](<https://github.com/sphireinc/mantis/blob/master/byte/bytes.go#L12>)

```go
func (b Bytes) Format(s fmt.State, verb rune)
```

Format bytes

<a name="Bytes.String"></a>
### func \(Bytes\) [String](<https://github.com/sphireinc/mantis/blob/master/byte/bytes.go#L49>)

```go
func (b Bytes) String() string
```



<a name="BytesUnit"></a>
## type [BytesUnit](<https://github.com/sphireinc/mantis/blob/master/byte/bytes_unit.go#L4>)

BytesUnit represents a unit of magnitude for Bytes in either SI \(base\-10\) or NIST \(base\-2\)

```go
type BytesUnit int64
```

<a name="BytesUnit.Convert"></a>
### func \(BytesUnit\) [Convert](<https://github.com/sphireinc/mantis/blob/master/byte/bytes_unit.go#L7>)

```go
func (u BytesUnit) Convert(b Bytes) float64
```

Convert a byte count into the \(possibly fractional\) count in this unit

<a name="BytesUnit.String"></a>
### func \(BytesUnit\) [String](<https://github.com/sphireinc/mantis/blob/master/byte/bytes_unit.go#L14>)

```go
func (u BytesUnit) String() string
```



<a name="ParseErr"></a>
## type [ParseErr](<https://github.com/sphireinc/mantis/blob/master/byte/parse.go#L11-L15>)

ParseErr returned invalid strings are parsed

```go
type ParseErr struct {
    Func string
    Str  string
    Err  error
}
```

<a name="ParseErr.Error"></a>
### func \(\*ParseErr\) [Error](<https://github.com/sphireinc/mantis/blob/master/byte/parse.go#L17>)

```go
func (e *ParseErr) Error() string
```



<a name="ParseErr.String"></a>
### func \(\*ParseErr\) [String](<https://github.com/sphireinc/mantis/blob/master/byte/parse.go#L21>)

```go
func (e *ParseErr) String() string
```



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)