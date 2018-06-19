package leftpad

import (
	"strings"
	"unicode/utf8"
)

var default_char = ' '

// Format is my comment
func Format(s string, size int) string {
	return FormatRune(s, size, default_char)
}

// FormatRune is my comment
func FormatRune(s string, size int, r rune) string {
	l := utf8.RuneCountInString(s)
	if l >= size {
		return s
	}
	out := strings.Repeat(string(r), size-l) + s
	return out
}
