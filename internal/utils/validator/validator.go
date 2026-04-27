package validator

import (
	"strconv"
	"strings"
)

func IsBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}

func IsInteger(s string) bool {
	num, err := strconv.Atoi(s)
	return err == nil && num >= 0
}

func StrToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}
