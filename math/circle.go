package math

import "math"

func areaOfCircle(radius float64) float64 {
	return math.Pi * square(radius)
}

func circumferenceOfCircle(radius float64) float64 {
	return 2 * math.Pi * radius
}
