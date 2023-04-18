package math

func fahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

func celsiusToFahrenheit(c float64) float64 {
	return (c * 9 / 5) + 32
}
