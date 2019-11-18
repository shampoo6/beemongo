package string_util

import "strings"

func IsEmpty(str string) bool {
	return strings.Trim(str, " ") == ""
}

func HasText(str string) bool {
	return !IsEmpty(str)
}
