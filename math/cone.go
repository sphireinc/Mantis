package math

import "math"

func volumeOfCone(radius, height float64) float64 {
	return (1.0 / 3.0) * math.Pi * square(radius) * height
}
