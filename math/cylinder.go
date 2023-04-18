package math

import "math"

func volumeOfCylinder(radius, height float64) float64 {
	return math.Pi * square(radius) * height
}

func surfaceAreaOfCylinder(radius, height float64) float64 {
	return 2*math.Pi*radius*height + 2*math.Pi*square(radius)
}
