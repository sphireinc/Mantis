package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCelsiusToFahrenheit(t *testing.T) {
	assert.Equal(t, CelsiusToFahrenheit(0), float32(32))
	assert.Equal(t, CelsiusToFahrenheit(100), float32(212))
}

func TestFahrenheitToCelsius(t *testing.T) {
	assert.Equal(t, FahrenheitToCelsius(32), float32(0))
	assert.Equal(t, FahrenheitToCelsius(212), float32(100))
}
