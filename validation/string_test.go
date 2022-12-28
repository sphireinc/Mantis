package validation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsAllUpper(t *testing.T) {
	assert.Equal(t, true, IsAllUpper("ABC"))
	assert.Equal(t, false, IsAllUpper(""))
	assert.Equal(t, false, IsAllUpper("abc"))
	assert.Equal(t, false, IsAllUpper("aBC"))
	assert.Equal(t, false, IsAllUpper("1BC"))
	assert.Equal(t, false, IsAllUpper("1bc"))
	assert.Equal(t, false, IsAllUpper("123"))
	assert.Equal(t, false, IsAllUpper("你好"))
	assert.Equal(t, false, IsAllUpper("A&"))
	assert.Equal(t, false, IsAllUpper("&@#$%^&*"))
}

func TestIsAllLower(t *testing.T) {
	assert.Equal(t, true, IsAllLower("abc"))
	assert.Equal(t, false, IsAllLower("ABC"))
	assert.Equal(t, false, IsAllLower(""))
	assert.Equal(t, false, IsAllLower("aBC"))
	assert.Equal(t, false, IsAllLower("1BC"))
	assert.Equal(t, false, IsAllLower("1bc"))
	assert.Equal(t, false, IsAllLower("123"))
	assert.Equal(t, false, IsAllLower("你好"))
	assert.Equal(t, false, IsAllLower("A&"))
	assert.Equal(t, false, IsAllLower("&@#$%^&*"))
}

func TestContainLower(t *testing.T) {
	assert.Equal(t, true, ContainLower("abc"))
	assert.Equal(t, true, ContainLower("aBC"))
	assert.Equal(t, true, ContainLower("1bc"))
	assert.Equal(t, true, ContainLower("a&"))

	assert.Equal(t, false, ContainLower("ABC"))
	assert.Equal(t, false, ContainLower(""))
	assert.Equal(t, false, ContainLower("1BC"))
	assert.Equal(t, false, ContainLower("123"))
	assert.Equal(t, false, ContainLower("你好"))
	assert.Equal(t, false, ContainLower("&@#$%^&*"))
}

func TestContainUpper(t *testing.T) {
	assert.Equal(t, true, ContainUpper("ABC"))
	assert.Equal(t, true, ContainUpper("aBC"))
	assert.Equal(t, true, ContainUpper("1BC"))
	assert.Equal(t, true, ContainUpper("A&"))

	assert.Equal(t, false, ContainUpper("abc"))
	assert.Equal(t, false, ContainUpper(""))
	assert.Equal(t, false, ContainUpper("1bc"))
	assert.Equal(t, false, ContainUpper("123"))
	assert.Equal(t, false, ContainUpper("你好"))
	assert.Equal(t, false, ContainUpper("&@#$%^&*"))
}

func TestContainLetter(t *testing.T) {
	assert.Equal(t, true, ContainLetter("ABC"))
	assert.Equal(t, true, ContainLetter("1Bc"))
	assert.Equal(t, true, ContainLetter("1ab"))
	assert.Equal(t, true, ContainLetter("A&"))

	assert.Equal(t, false, ContainLetter(""))
	assert.Equal(t, false, ContainLetter("123"))
	assert.Equal(t, false, ContainLetter("你好"))
	assert.Equal(t, false, ContainLetter("&@#$%^&*"))
}

func TestIsNumberStr(t *testing.T) {
	assert.Equal(t, true, IsNumberStr("3."))
	assert.Equal(t, true, IsNumberStr("+3."))
	assert.Equal(t, true, IsNumberStr("-3."))
	assert.Equal(t, true, IsNumberStr("+3e2"))
	assert.Equal(t, false, IsNumberStr("abc"))
}

func TestIsFloatStr(t *testing.T) {
	assert.Equal(t, true, IsFloatStr("3."))
	assert.Equal(t, true, IsFloatStr("+3."))
	assert.Equal(t, true, IsFloatStr("-3."))
	assert.Equal(t, true, IsFloatStr("12"))
	assert.Equal(t, false, IsFloatStr("abc"))
}

func TestIsEmptyString(t *testing.T) {
	assert.Equal(t, true, IsEmptyString(""))
	assert.Equal(t, false, IsEmptyString("111"))
	assert.Equal(t, false, IsEmptyString(" "))
	assert.Equal(t, false, IsEmptyString("\t"))
}

func TestIsAlpha(t *testing.T) {
	assert.Equal(t, true, IsAlpha("abc"))
	assert.Equal(t, false, IsAlpha("111"))
	assert.Equal(t, false, IsAlpha(" "))
	assert.Equal(t, false, IsAlpha("\t"))
	assert.Equal(t, false, IsAlpha(""))
}
