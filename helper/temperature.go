package helper

func CelsiusToFahrenheit(degrees float32) float32 {
	return (degrees * 1.8) + 32
}

func FahrenheitToCelsius(degrees float32) float32 {
	return (degrees - 32) / 1.8
}
