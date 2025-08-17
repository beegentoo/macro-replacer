package mare

import "strings"

func Upper(str string) string {
	return strings.ToUpper(str)
}

func Lower(str string) string {
	return strings.ToLower(str)
}

func Firstletter(str string) string {
	return str[0:1]
}
