package byte

import (
	"fmt"
	"strconv"
)

// Bytes is a formattable byte count
type Bytes int64

// Format bytes
func (b Bytes) Format(s fmt.State, verb rune) {
	switch verb {
	case 'd':
		_, _ = fmt.Fprintf(s, "%d B", int64(b))
	case 's':
		fallthrough
	case 'v':
		bCopy := b
		if bCopy < 0 {
			_, _ = s.Write([]byte{'-'})
			bCopy = b * -1
		}
		thresholds := bytesUnitThresholds
		if s.Flag('+') {
			thresholds = bytesSIUnitThresholds
		}
		for _, t := range thresholds {
			if bCopy < t.less {
				f := t.unit.Convert(bCopy)
				if prec, ok := s.Precision(); ok {
					_, _ = fmt.Fprintf(s, fmt.Sprintf("%%.%df %%s", prec), f, t.unit)
				} else {
					_, _ = fmt.Fprintf(s, "%s %s", strconv.FormatFloat(f, 'f', -1, 64), t.unit)
				}
				return
			}
		}
	}
}

func (b Bytes) String() string {
	return fmt.Sprintf("%s", b)
}
