package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  I'm a Little Teapot Short and Stout  ",
			expected: []string{"i'm", "a", "little", "teapot", "short", "and", "stout"},
		},
		{
			input:    "HELLO WORLD!",
			expected: []string{"hello", "world!"},
		},
		{
			input:    "1234 @#$%^&",
			expected: []string{"1234", "@#$%^&"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("For input '%s', expected length was %d, but got %d", c.input, len(c.expected), len(actual))
			continue
		}
		for i := range actual {
			word := actual[i]
			if word != c.expected[i] {
				t.Fatalf("For input '%s', expected '%s' at index %d, but got '%s'", c.input, c.expected[i], i, word)
			}
		}
	}
}
