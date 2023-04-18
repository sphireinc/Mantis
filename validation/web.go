package validation

import (
	"net"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

// IsPort determines if a string is a port
func IsPort(str string) bool {
	i, err := strconv.ParseInt(str, 10, 64)
	if err == nil && i > 0 && i < 65536 {
		return true
	}
	return false
}

// IsDNS determines if a string is a valid DNS entry
func IsDNS(str string) bool {
	return regexp.MustCompile(`^[a-zA-Z]([a-zA-Z0-9\-]+[\.]?)*[a-zA-Z0-9]$`).MatchString(str)
}

// IsEmail determines if a string is a valid email
func IsEmail(str string) bool {
	return regexp.MustCompile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`).MatchString(str)
}

// IsIP determines if a string is a valid IP address
func IsIP(str string) bool {
	return net.ParseIP(str) != nil
}

// IsIPV4 determines if a string is a valid IPv4 address
func IsIPV4(str string) bool {
	return net.ParseIP(str) != nil && strings.Contains(str, ".")
}

// IsIPV6 determines if a string is a valid IPv6 character
func IsIPV6(str string) bool {
	return net.ParseIP(str) != nil && strings.Contains(str, ":")
}

// IsURL determines if a string is a valid URL
func IsURL(str string) bool {
	if str == "" || len(str) >= 2083 || len(str) <= 3 || strings.HasPrefix(str, ".") {
		return false
	}
	u, err := url.Parse(str)
	if err != nil {
		return false
	}
	if strings.HasPrefix(u.Host, ".") {
		return false
	}
	if u.Host == "" && (u.Path != "" && !strings.Contains(u.Path, ".")) {
		return false
	}

	matcher := regexp.MustCompile(`^((ftp|http|https?):\/\/)?(\S+(:\S*)?@)?((([1-9]\d?|1\d\d|2[01]\d|22[0-3])(\.(1?\d{1,2}|2[0-4]\d|25[0-5])){2}(?:\.([0-9]\d?|1\d\d|2[0-4]\d|25[0-4]))|(([a-zA-Z0-9]+([-\.][a-zA-Z0-9]+)*)|((www\.)?))?(([a-z\x{00a1}-\x{ffff}0-9]+-?-?)*[a-z\x{00a1}-\x{ffff}0-9]+)(?:\.([a-z\x{00a1}-\x{ffff}]{2,}))?))(:(\d{1,5}))?((\/|\?|#)[^\s]*)?$`)
	return matcher.MatchString(str)
}
