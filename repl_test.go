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
			input:    "  justOneWord  ",
			expected: []string{"justoneword"},
		},
		{
			input:    "      ",
			expected: []string{},
		},
		{
			input:    "This is A robbery",
			expected: []string{"this", "is", "a", "robbery"},
		},
		{
			input:    "YES WE CAN     ",
			expected: []string{"yes", "we", "can"},
		},
		{
			input:    "    HOLD Up, someThinG aIn'T RighT",
			expected: []string{"hold", "up,", "something", "ain't", "right"},
		},
		{
			input:    " !! $%, 75789 (==) =£/ ",
			expected: []string{"!!", "$%,", "75789", "(==)", "=£/"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		if len(actual) != len(c.expected) {
			t.Errorf("Lengths don't match for input '%s': expected '%d', got '%d'",
				c.input, len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Words don't match for input '%s': expected '%s', got '%s'",
					c.input, expectedWord, word)
			}
		}
	}
}
