package dsutils

import "slices"

// note - not necessarily correct for non-ASCII strings
func ReverseString(str string) string {
	runes := []rune(str)
	slices.Reverse(runes)
	return string(runes)
}
