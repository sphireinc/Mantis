package geo

import "math"

func decimalToDMS(decimal float64) (int, int, float64) {
	degrees := int(math.Floor(math.Abs(decimal)))
	minutesFloat := (math.Abs(decimal) - float64(degrees)) * 60
	minutes := int(math.Floor(minutesFloat))
	seconds := (minutesFloat - float64(minutes)) * 60

	if decimal < 0 {
		degrees = -degrees
	}

	return degrees, minutes, seconds
}

func dmsToDecimal(degrees, minutes int, seconds float64) float64 {
	sign := 1.0
	if degrees < 0 {
		sign = -1.0
	}

	decimal := float64(degrees) + float64(minutes)/60 + seconds/3600

	return sign * decimal
}
