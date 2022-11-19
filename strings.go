package helper

import "strings"

// Contains checks if a string contains a substring case insensitive
func Contains(a string, b string) bool {
	return strings.Contains(strings.ToLower(a), strings.ToLower(b))
}
