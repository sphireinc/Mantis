package byte

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// ParseErr returned invalid strings are parsed
type ParseErr struct {
	Func string
	Str  string
	Err  error
}

func (e *ParseErr) Error() string {
	return e.String()
}

func (e *ParseErr) String() string {
	return "bytes." + e.Func + ": " + "parsing " + strconv.Quote(e.Str) + ": " + e.Err.Error()
}

// ParseBytes parses the given string with byte units into a Byte count
func ParseBytes(s string) (Bytes, error) {
	regex := `(?i)(?P<value>[0-9.+eE+-]+)\s*(?P<unit>[a-z]?(?:i?b)?)?`
	gs := regexp.MustCompile(regex).FindStringSubmatch(s)
	if len(gs) == 0 {
		return 0, &ParseErr{
			Func: "ParseBytes",
			Str:  s,
			Err:  errors.New("bad string format"),
		}
	}
	vs := strings.ToLower(gs[1])
	us := gs[2]

	unit, ok := strToByteUnit[strings.ToUpper(strings.TrimSpace(us))]
	if !ok {
		return 0, &ParseErr{
			Func: "ParseBytes",
			Str:  s,
			Err:  errors.New("unknown bytes unit"),
		}
	}

	vs = strings.Replace(vs, ",", "", -1)
	if strings.Contains(vs, ".") || strings.Contains(vs, "e") {
		f, err := strconv.ParseFloat(vs, 64)
		if err != nil {
			return 0, &ParseErr{
				Func: "ParseBytes",
				Str:  s,
				Err:  err,
			}
		}
		return Bytes(f * float64(unit)), nil
	}
	i, err := strconv.ParseInt(vs, 10, 64)
	if err != nil {
		return 0, &ParseErr{
			Func: "ParseBytes",
			Str:  s,
			Err:  err,
		}
	}
	return Bytes(i * int64(unit)), nil
}
