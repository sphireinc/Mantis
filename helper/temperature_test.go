package helper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCelsiusToFahrenheit(t *testing.T) {
	assert.Equal(t, CelsiusToFahrenheit(0), 32)
	assert.Equal(t, CelsiusToFahrenheit(100), 200)
}

func TestFahrenheitToCelsius(t *testing.T) {
	assert.Equal(t, FahrenheitToCelsius(32), 0)
	assert.Equal(t, FahrenheitToCelsius(200), 100)
}
