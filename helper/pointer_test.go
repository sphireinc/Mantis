package helper

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
