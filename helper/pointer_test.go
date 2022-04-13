package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToPtr(t *testing.T) {
	var val = struct {
		Name string
	}{
		Name: "Mantis",
	}

	assert.Equal(t, &val, ToPtr(val))
}

func TestFromPtr(t *testing.T) {
	var val = struct {
		Name string
	}{
		Name: "Mantis",
	}

	assert.Equal(t, val, FromPtr(&val))
}
