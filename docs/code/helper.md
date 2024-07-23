<!-- Code generated by gomarkdoc. DO NOT EDIT -->

[![Go Report Card](https://goreportcard.com/badge/github.com/sphireinc/mantis)](https://goreportcard.com/report/github.com/sphireinc/mantis)
[![Go Reference](https://pkg.go.dev/badge/github.com/sphireinc/mantis.svg)](https://pkg.go.dev/github.com/sphireinc/mantis)



# helper

```go
import "github.com/sphireinc/mantis/helper"
```

## Index

- [func AtoiWithDefault\(value string, defaultValue int\) int](<#AtoiWithDefault>)
- [func Average\[T IntFloat\]\(args ...T\) float64](<#Average>)
- [func CelsiusToFahrenheit\(degrees float32\) float32](<#CelsiusToFahrenheit>)
- [func Default\[T comparable\]\(originalVal T, defaultVal T\) T](<#Default>)
- [func DeferFileClose\(file \*os.File\)](<#DeferFileClose>)
- [func FahrenheitToCelsius\(degrees float32\) float32](<#FahrenheitToCelsius>)
- [func FromPtr\[T any\]\(original \*T\) T](<#FromPtr>)
- [func Reverse\(s string\) string](<#Reverse>)
- [func StringToBool\(boolean string\) bool](<#StringToBool>)
- [func ToPtr\[T any\]\(original T\) \*T](<#ToPtr>)
- [type IntFloat](<#IntFloat>)


<a name="AtoiWithDefault"></a>
## func [AtoiWithDefault](<https://github.com/sphireinc/mantis/blob/master/helper/helper.go#L32>)

```go
func AtoiWithDefault(value string, defaultValue int) int
```

AtoiWithDefault same as strconv.Atoi except only returns the value or a default value if nil

<a name="Average"></a>
## func [Average](<https://github.com/sphireinc/mantis/blob/master/helper/math.go#L9>)

```go
func Average[T IntFloat](args ...T) float64
```

Average calculates the average value of a set of numbers

<a name="CelsiusToFahrenheit"></a>
## func [CelsiusToFahrenheit](<https://github.com/sphireinc/mantis/blob/master/helper/temperature.go#L4>)

```go
func CelsiusToFahrenheit(degrees float32) float32
```

CelsiusToFahrenheit converts C to F

<a name="Default"></a>
## func [Default](<https://github.com/sphireinc/mantis/blob/master/helper/helper.go#L42>)

```go
func Default[T comparable](originalVal T, defaultVal T) T
```

Default returns the defaultVal given if originalVal is empty/nil

<a name="DeferFileClose"></a>
## func [DeferFileClose](<https://github.com/sphireinc/mantis/blob/master/helper/helper.go#L18>)

```go
func DeferFileClose(file *os.File)
```

DeferFileClose prevents non\-closure file closing error

<a name="FahrenheitToCelsius"></a>
## func [FahrenheitToCelsius](<https://github.com/sphireinc/mantis/blob/master/helper/temperature.go#L9>)

```go
func FahrenheitToCelsius(degrees float32) float32
```

FahrenheitToCelsius converts F to C

<a name="FromPtr"></a>
## func [FromPtr](<https://github.com/sphireinc/mantis/blob/master/helper/pointer.go#L9>)

```go
func FromPtr[T any](original *T) T
```

FromPtr returns a pointer value for the value passed in.

<a name="Reverse"></a>
## func [Reverse](<https://github.com/sphireinc/mantis/blob/master/helper/helper.go#L9>)

```go
func Reverse(s string) string
```

Reverse a string

<a name="StringToBool"></a>
## func [StringToBool](<https://github.com/sphireinc/mantis/blob/master/helper/helper.go#L23>)

```go
func StringToBool(boolean string) bool
```

StringToBool same as strconv.ParseBool except hides the error \(returns false\)

<a name="ToPtr"></a>
## func [ToPtr](<https://github.com/sphireinc/mantis/blob/master/helper/pointer.go#L4>)

```go
func ToPtr[T any](original T) *T
```

ToPtr returns a pointer value for the value passed in.

<a name="IntFloat"></a>
## type [IntFloat](<https://github.com/sphireinc/mantis/blob/master/helper/math.go#L4-L6>)

IntFloat is an interface for numeric types

```go
type IntFloat interface {
    float32 | float64 | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}
```

Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)