package uuid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	n := New()
	assert.NotEmpty(t, n.UUID.String())
}
