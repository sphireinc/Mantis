package validation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPort(t *testing.T) {
	assert.Equal(t, true, IsPort("1"))
	assert.Equal(t, true, IsPort("65535"))
	assert.Equal(t, false, IsPort("abc"))
	assert.Equal(t, false, IsPort("123abc"))
	assert.Equal(t, false, IsPort(""))
	assert.Equal(t, false, IsPort("0"))
	assert.Equal(t, false, IsPort("-1"))
	assert.Equal(t, false, IsPort("65536"))
}

func TestIsIp(t *testing.T) {
	assert.Equal(t, true, IsIp("127.0.0.1"))
	assert.Equal(t, true, IsIp("::0:0:0:0:0:0:1"))
	assert.Equal(t, false, IsIp("127.0.0"))
	assert.Equal(t, false, IsIp("127"))
}

func TestIsIpV4(t *testing.T) {
	assert.Equal(t, true, IsIPV4("127.0.0.1"))
	assert.Equal(t, false, IsIPV4("::0:0:0:0:0:0:1"))
}

func TestIsIpV6(t *testing.T) {
	assert.Equal(t, false, IsIPV6("127.0.0.1"))
	assert.Equal(t, true, IsIPV6("::0:0:0:0:0:0:1"))
}

func TestIsUrl(t *testing.T) {
	assert.Equal(t, true, IsURL("http://abc.com"))
	assert.Equal(t, true, IsURL("https://abc.com"))
	assert.Equal(t, true, IsURL("abc.com"))
	assert.Equal(t, true, IsURL("a.b.com"))
	assert.Equal(t, false, IsURL("abc"))
}

func TestIsDns(t *testing.T) {
	assert.Equal(t, true, IsDNS("abc.com"))
	assert.Equal(t, false, IsDNS("a.b.com"))
	assert.Equal(t, false, IsDNS("http://abc.com"))
}

func TestIsEmail(t *testing.T) {
	assert.Equal(t, true, IsEmail("abc@xyz.com"))
	assert.Equal(t, false, IsEmail("a.b@@com"))
}
