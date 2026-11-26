package main

import "testing"

func TestCleanInput(t *testing.T) {
	type testCase struct {
		input    string
		expected []string
	}

	cases := []testCase{
		{
			input:    "  ",
			expected: []string{},
		},
		{
			input:    "  hello  ",
			expected: []string{"hello"},
		},
		{
			input:    "   hello   world         ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		got := cleanInput(c.input)
		if len(got) != len(c.expected) {
			t.Errorf("length is wrong, got %d, want %d", len(got), len(c.expected))
		}
		for i := range got {
			gotWord := got[i]
			expectedWord := c.expected[i]

			if gotWord != expectedWord {
				t.Errorf("words not equal, got %s, want %s", gotWord, expectedWord)
			}
		}
	}
}
