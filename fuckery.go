package fuckery

import (
	"strings"
)

const (
	StrikeRune    rune = '\u0336'
	UnderlineRune rune = '\u0332'
)

// appendRune appends a rune after each non-newline rune in a string.
func appendRune(s string, r rune) string {
	var b strings.Builder
	b.Grow(len(s) * 2)
	for _, c := range s {
		b.WriteRune(c)
		if c != '\n' {
			b.WriteRune(r)
		}
	}
	return b.String()
}

// Strike modifies a unicode string to insert a diacritical strikethrough after
// non newline runes.
func Strike(s string) string {
	return appendRune(s, rune(StrikeRune))
}

// Underline modifies a unicode string to insert a diacritical underline after
// non newline runes.
func Underline(text string) string {
	return appendRune(text, rune(UnderlineRune))
}
