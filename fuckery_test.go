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
		{"a", "a̶"},
		{"A", "A̶"},
		{"Hello World", "H̶e̶l̶l̶o̶ ̶W̶o̶r̶l̶d̶"},
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
		{"a", "a̲"},
		{"A", "A̲"},
		{"Hello World", "H̲e̲l̲l̲o̲ ̲W̲o̲r̲l̲d̲"},
	}
	for _, test := range tests {
		if got := Underline(test.input); got != test.want {
			t.Errorf("Underline(%q) = %v", test.input, got)
		}
	}
}
