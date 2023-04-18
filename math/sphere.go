package math

import "math"

func volumeOfSphere(radius float64) float64 {
	return (4.0 / 3.0) * math.Pi * math.Pow(radius, 3)
}

func surfaceAreaOfSphere(radius float64) float64 {
	return 4 * math.Pi * square(radius)
}
