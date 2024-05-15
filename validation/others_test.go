package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsJSON(t *testing.T) {
	assert.Equal(t, true, IsJSON("{}"))
	assert.Equal(t, true, IsJSON("{\"name\": \"test\"}"))
	assert.Equal(t, true, IsJSON("[]"))
	assert.Equal(t, true, IsJSON("123"))

	assert.Equal(t, false, IsJSON(""))
	assert.Equal(t, false, IsJSON("abc"))
	assert.Equal(t, false, IsJSON("你好"))
	assert.Equal(t, false, IsJSON("&@#$%^&*"))
}

func TestIsRegexMatch(t *testing.T) {
	assert.Equal(t, true, IsRegexMatch("abc", `^[a-zA-Z]+$`))
	assert.Equal(t, false, IsRegexMatch("1ab", `^[a-zA-Z]+$`))
	assert.Equal(t, false, IsRegexMatch("", `^[a-zA-Z]+$`))
}

func TestIsBase64(t *testing.T) {
	assert.Equal(t, true, IsBase64("aGVsbG8="))
	assert.Equal(t, false, IsBase64("123456"))
}
