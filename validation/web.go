package validation

import (
	"net"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

func IsPort(str string) bool {
	i, err := strconv.ParseInt(str, 10, 64)
	if err == nil && i > 0 && i < 65536 {
		return true
	}
	return false
}

func IsDns(str string) bool {
	return regexp.MustCompile(`^[a-zA-Z]([a-zA-Z0-9\-]+[\.]?)*[a-zA-Z0-9]$`).MatchString(str)
}

func IsEmail(str string) bool {
	return regexp.MustCompile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`).MatchString(str)
}

func IsIp(str string) bool {
	return net.ParseIP(str) != nil
}

func IsIpV4(str string) bool {
	return net.ParseIP(str) != nil && strings.Contains(str, ".")
}

func IsIpV6(str string) bool {
	return net.ParseIP(str) != nil && strings.Contains(str, ":")
}

func IsUrl(str string) bool {
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
