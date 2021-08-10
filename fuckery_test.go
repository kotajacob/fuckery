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
