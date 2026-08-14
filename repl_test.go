package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "HELLO",
			expected: []string{"hello"},
		},
		{
			input:    "\tmixed\n whitespace\r\n here  ",
			expected: []string{"mixed", "whitespace", "here"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "     ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput(%q): expected %d words, got %d (%v)",
				c.input, len(c.expected), len(actual), actual)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%q): word %d: expected %q, got %q",
					c.input, i, expectedWord, word)
			}
		}
	}
}
