package helper

// CelsiusToFahrenheit converts C to F
func CelsiusToFahrenheit(degrees float32) float32 {
	return (degrees * 1.8) + 32
}

// FahrenheitToCelsius converts F to C
func FahrenheitToCelsius(degrees float32) float32 {
	return (degrees - 32) / 1.8
}
