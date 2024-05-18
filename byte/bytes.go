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
		// Format as a decimal integer with "B" suffix
		_, _ = fmt.Fprintf(s, "%d B", int64(b))

	case 's', 'v':
		bCopy := b
		// Handle negative values
		if bCopy < 0 {
			_, _ = s.Write([]byte{'-'})
			bCopy = -bCopy
		}

		// Choose the appropriate unit thresholds
		thresholds := bytesUnitThresholds
		if s.Flag('+') {
			thresholds = bytesSIUnitThresholds
		}

		// Find the appropriate unit threshold and format the output
		for _, t := range thresholds {
			if bCopy < t.less {
				f := t.unit.Convert(bCopy)
				if prec, ok := s.Precision(); ok {
					// Use the provided precision
					_, _ = fmt.Fprintf(s, fmt.Sprintf("%%.%df %%s", prec), f, t.unit)
				} else {
					// Default formatting without specific precision
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
