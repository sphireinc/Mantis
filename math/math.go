package math

func square(x float64) float64 {
	return x * x
}

func isEven(x int) bool {
	return x%2 == 0
}

func isOdd(x int) bool {
	return x%2 != 0
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

// gcd calculates the greatest common divisor
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

// lcm calculates the least common multiple
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func firstNPrimes(n int) []int {
	primes := []int{}
	i := 2
	for len(primes) < n {
		if isPrime(i) {
			primes = append(primes, i)
		}
		i++
	}
	return primes
}
