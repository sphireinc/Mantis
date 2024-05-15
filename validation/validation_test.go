package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsChinese(t *testing.T) {
	assert.Equal(t, true, ContainsChinese("你好"))
	assert.Equal(t, true, ContainsChinese("你好hello"))
	assert.Equal(t, false, ContainsChinese("hello"))
}
