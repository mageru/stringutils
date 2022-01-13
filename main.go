package stringutils

import (
	"strings"
)

func Upper(str string) string {
	var upper_str = strings.ToUpper(str)
	return upper_str
}

func Lower(str string) string {
	var lower_str = strings.ToLower(str)
	return lower_str
}
