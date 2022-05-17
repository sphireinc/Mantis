package helper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAverage1(t *testing.T) {
	assert.Equal(t, float64(40), Average(32, 34, 54))
	assert.Equal(t, float64(1), Average(1))
	assert.Equal(t, float64(2), Average(1, 2, 3))
}
