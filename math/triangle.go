package math

import "math"

func areaOfTriangle(base, height float64) float64 {
	return 0.5 * base * height
}

func perimeterOfTriangle(a, b, c float64) float64 {
	return a + b + c
}

func hypotenuseLength(a, b float64) float64 {
	return math.Sqrt(square(a) + square(b))
}
