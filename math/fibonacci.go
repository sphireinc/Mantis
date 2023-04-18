package math

import "math"

func isFibonacci(n int) bool {
	a := 5*square(float64(n)) + 4
	b := 5*square(float64(n)) - 4

	return isPerfectSquare(int(a)) || isPerfectSquare(int(b))
}

func isPerfectSquare(x int) bool {
	return int(square(math.Sqrt(float64(x)))) == x
}

// fibonacci calculates the nth Fibonacci number
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
