package iteration

import "strings"

// Repeat takes a string and an integer, and return a new string consisting of the input string repeated the specified number of times.
func Repeat(character string, count int) string {
	var repeated strings.Builder
	for i := 0; i < count; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
