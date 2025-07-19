package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello",
			expected: []string{"hello"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    " HELLO WORLD again   and  again",
			expected: []string{"hello", "world", "again", "and", "again"},
		},
	}
	for i, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("\nLengths do not match in case %d:\n\tExpected: %d\n\tActual: %d", i, len(c.expected), len(actual))
		}
		for j, word := range actual {
			if word != c.expected[j] {
				t.Errorf("\nStrings do not match in case %d item %d:\n\tExpected: %s\n\tActual: %s", i, j, word, c.expected[j])
			}
		}
	}
}
