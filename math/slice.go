package math

import "math"

func sumIntSlice(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func meanFloatSlice(numbers []float64) float64 {
	sum := 0.0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}

func maxIntSlice(numbers []int) int {
	max := numbers[0]
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}

func minIntSlice(numbers []int) int {
	min := numbers[0]
	for _, number := range numbers {
		if number < min {
			min = number
		}
	}
	return min
}

// calculate the dot product of two float64 slices
func dotProduct(a, b []float64) float64 {
	if len(a) != len(b) {
		panic("Vectors must have the same length")
	}

	dot := 0.0
	for i := range a {
		dot += a[i] * b[i]
	}
	return dot
}

// calculate the Euclidean distance between two float64 slices
func euclideanDistance(a, b []float64) float64 {
	if len(a) != len(b) {
		panic("Vectors must have the same length")
	}

	sum := 0.0
	for i := range a {
		sum += square(a[i] - b[i])
	}
	return math.Sqrt(sum)
}

func uniqueIntSlice(numbers []int) []int {
	unique := make([]int, 0)
	seen := make(map[int]bool)

	for _, number := range numbers {
		if !seen[number] {
			seen[number] = true
			unique = append(unique, number)
		}
	}
	return unique
}
