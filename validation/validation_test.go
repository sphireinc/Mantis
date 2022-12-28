package validation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainsChinese(t *testing.T) {
	assert.Equal(t, true, ContainsChinese("你好"))
	assert.Equal(t, true, ContainsChinese("你好hello"))
	assert.Equal(t, false, ContainsChinese("hello"))
}
