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
func Underline(s string) string {
	return appendRune(s, rune(UnderlineRune))
}

// Offset represents a range of runes that can be shifted by a certain number
// of unicode places. Use a starting character, ending character, and a
// replacement for the starting character at the new shifted position.
type Offset struct {
	start       rune
	end         rune
	replacement rune
}

// offsetRune modified a unicode string to replace characters based on slice of
// Offsets.
func offsetRune(s string, offsets []Offset) string {
	var b strings.Builder
	b.Grow(len(s))
	for _, c := range s {
		r := c
		// check if the rune exists within this offset
		for _, o := range offsets {
			if c >= o.start && c <= o.end {
				r = c + (o.replacement - o.start)
			}
		}
		b.WriteRune(r)
	}
	return b.String()
}

// BoldSans modifies a unicode string to replace ASCII characters with their
// 𝗯𝗼𝗹𝗱 𝘀𝗮𝗻𝘀 version.
func BoldSans(s string) string {
	offsets := []Offset{
		{'A', 'Z', '𝗔'},
		{'a', 'z', '𝗮'},
		{'0', '9', '𝟬'},
	}
	return offsetRune(s, offsets)
}

// BoldSerif modifies a unicode string to replace ASCII characters with their
// 𝐛𝐨𝐥𝐝 𝐬𝐞𝐫𝐢𝐟 version.
func BoldSerif(s string) string {
	offsets := []Offset{
		{'A', 'Z', '𝐀'},
		{'a', 'z', '𝐚'},
		{'0', '9', '𝟎'},
	}
	return offsetRune(s, offsets)
}

// ItalicSans modifies a unicode string to replace ASCII characters with their
// 𝘪𝘵𝘢𝘭𝘪𝘤 𝘴𝘢𝘯𝘴 version.
func ItalicSans(s string) string {
	offsets := []Offset{
		{'A', 'Z', '𝘈'},
		{'a', 'z', '𝘢'},
	}
	return offsetRune(s, offsets)
}

// ItalicSerif modifies a unicode string to replace ASCII characters with their
// 𝑖𝑡𝑎𝑙𝑖𝑐 𝑠𝑒𝑟𝑖𝑓 version.
func ItalicSerif(s string) string {
	offsets := []Offset{
		{'A', 'Z', '𝐴'},
		{'a', 'z', '𝑎'},
	}
	return offsetRune(s, offsets)
}

// BoldItalicSans modifies a unicode string to replace ASCII characters with
// their 𝙗𝙤𝙡𝙙 𝙞𝙩𝙖𝙡𝙞𝙘 𝙨𝙖𝙣𝙨 version.
func BoldItalicSans(s string) string {
	offsets := []Offset{
		{'A', 'Z', '𝘼'},
		{'a', 'z', '𝙖'},
		{'0', '9', '𝟬'},
	}
	return offsetRune(s, offsets)
}

// BoldItalicSerif modifies a unicode string to replace ASCII characters with
// their 𝒃𝒐𝒍𝒅 𝒊𝒕𝒂𝒍𝒊𝒄 𝒔𝒆𝒓𝒊𝒇 version.
func BoldItalicSerif(s string) string {
	offsets := []Offset{
		{'A', 'Z', '𝑨'},
		{'a', 'z', '𝒂'},
		{'0', '9', '𝟎'},
	}
	return offsetRune(s, offsets)
}
