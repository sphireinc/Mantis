package validation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsStrongPassword(t *testing.T) {
	isStrong, complexity := IsStrongPassword("abc", 3)
	assert.Equal(t, false, isStrong)
	assert.Equal(t, 3, complexity)

	isStrong, complexity = IsStrongPassword("abc123", 6)
	assert.Equal(t, false, isStrong)
	assert.Equal(t, 12, complexity)

	isStrong, complexity = IsStrongPassword("abcABC", 6)
	assert.Equal(t, false, isStrong)
	assert.Equal(t, 9, complexity)

	isStrong, complexity = IsStrongPassword("abc123@#$", 9)
	assert.Equal(t, true, isStrong)
	assert.Equal(t, 27, complexity)

	isStrong, complexity = IsStrongPassword("abcABC123@#$", 16)
	assert.Equal(t, false, isStrong)
	assert.Equal(t, 0, complexity)

	isStrong, complexity = IsStrongPassword("abcABC123@#$Ed5%", 16)
	assert.Equal(t, true, isStrong)
	assert.Equal(t, 44, complexity)

	isStrong, complexity = IsStrongPassword("abcABC123@#$", 12)
	assert.Equal(t, true, isStrong)
	assert.Equal(t, 33, complexity)

	isStrong, complexity = IsStrongPassword("abcABC123@#$", 10)
	assert.Equal(t, true, isStrong)
	assert.Equal(t, 33, complexity)
}

func TestIsWeakPassword(t *testing.T) {
	assert.Equal(t, true, IsWeakPassword("abc"))
	assert.Equal(t, true, IsWeakPassword("123"))
	assert.Equal(t, true, IsWeakPassword("abc123"))
	assert.Equal(t, true, IsWeakPassword("abcABC123"))
	assert.Equal(t, false, IsWeakPassword("abc123@#$"))
}
