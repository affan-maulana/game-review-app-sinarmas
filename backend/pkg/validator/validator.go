package validator

import "strings"

func TrimString(s string) string {
	return strings.TrimSpace(s)
}

func IsEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}
