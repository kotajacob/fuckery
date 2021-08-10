package fuckery

import (
	"testing"
)

func TestStrike(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"", ""},
		{"\n", "\n"},
		{"a", "a\u0336"},
		{"A", "A\u0336"},
		{"Hello, 世界",
			"H\u0336e\u0336l\u0336l\u0336o\u0336,\u0336 \u0336世\u0336界\u0336"},
	}
	for _, test := range tests {
		if got := Strike(test.input); got != test.want {
			t.Errorf("Strike(%q) = %v", test.input, got)
		}
	}
}

func TestUnderline(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"", ""},
		{"\n", "\n"},
		{"a", "a\u0332"},
		{"A", "A\u0332"},
		{"Hello, 世界",
			"H\u0332e\u0332l\u0332l\u0332o\u0332,\u0332 \u0332世\u0332界\u0332"},
	}
	for _, test := range tests {
		if got := Underline(test.input); got != test.want {
			t.Errorf("Underline(%q) = %v", test.input, got)
		}
	}
}

func TestBoldSans(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"", ""},
		{"\n", "\n"},
		{"a", "𝗮"},
		{"A", "𝗔"},
		{"Hello, 世界", "𝗛𝗲𝗹𝗹𝗼, 世界"},
	}
	for _, test := range tests {
		if got := BoldSans(test.input); got != test.want {
			t.Errorf("BoldSans(%q) = %v", test.input, got)
		}
	}
}

func TestBoldSerif(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"", ""},
		{"\n", "\n"},
		{"a", "𝐚"},
		{"A", "𝐀"},
		{"Hello, 世界", "𝐇𝐞𝐥𝐥𝐨, 世界"},
	}
	for _, test := range tests {
		if got := BoldSerif(test.input); got != test.want {
			t.Errorf("BoldSerif(%q) = %v", test.input, got)
		}
	}
}

func TestItalicSans(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"", ""},
		{"\n", "\n"},
		{"a", "𝘢"},
		{"A", "𝘈"},
		{"Hello, 世界", "𝘏𝘦𝘭𝘭𝘰, 世界"},
	}
	for _, test := range tests {
		if got := ItalicSans(test.input); got != test.want {
			t.Errorf("ItalicSans(%q) = %v", test.input, got)
		}
	}
}

func TestItalicSerif(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"", ""},
		{"\n", "\n"},
		{"a", "𝑎"},
		{"A", "𝐴"},
		{"Hello, 世界", "𝐻𝑒𝑙𝑙𝑜, 世界"},
	}
	for _, test := range tests {
		if got := ItalicSerif(test.input); got != test.want {
			t.Errorf("ItalicSerif(%q) = %v", test.input, got)
		}
	}
}

func TestBoldItalicSans(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"", ""},
		{"\n", "\n"},
		{"a", "𝙖"},
		{"A", "𝘼"},
		{"Hello, 世界", "𝙃𝙚𝙡𝙡𝙤, 世界"},
	}
	for _, test := range tests {
		if got := BoldItalicSans(test.input); got != test.want {
			t.Errorf("BoldItalicSans(%q) = %v", test.input, got)
		}
	}
}

func TestBoldItalicSerif(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"", ""},
		{"\n", "\n"},
		{"a", "𝒂"},
		{"A", "𝑨"},
		{"Hello, 世界", "𝑯𝒆𝒍𝒍𝒐, 世界"},
	}
	for _, test := range tests {
		if got := BoldItalicSerif(test.input); got != test.want {
			t.Errorf("BoldItalicSerif(%q) = %v", test.input, got)
		}
	}
}
