package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected []string
	}{{
		name:     "single word",
		input:    "pikachu",
		expected: []string{"pikachu"},
	}, {
		name:     "empty input",
		input:    "",
		expected: []string{},
	}, {
		name:     "whitespace only",
		input:    "    \t     ",
		expected: []string{},
	}, {
		name:     "tab between words",
		input:    "hello\tworld",
		expected: []string{"hello", "world"},
	}, {
		name:     "newline between words",
		input:    "hello\nworld",
		expected: []string{"hello", "world"},
	}, {
		name:     "mixed case",
		input:    "piKaCHu",
		expected: []string{"pikachu"},
	}, {
		name:     "multiple words with extra whitespace",
		input:    "  hello  world  ",
		expected: []string{"hello", "world"},
	}, {
		name:     "multiple words with uppercase letters and words",
		input:    "Charmander Bulbasaur PIKACHU",
		expected: []string{"charmander", "bulbasaur", "pikachu"},
	},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := cleanInput(c.input)
			// Check the length of the actual slice
			//  if they don't match, use t.Errorf and continue to the next case
			if len(actual) != len(c.expected) {
				t.Errorf("incorrect return length. Expected: %d, Got: %d", len(c.expected), len(actual))
				return
			}
			for i := range actual {
				word := actual[i]
				expectedWord := c.expected[i]
				if word != expectedWord {
					t.Errorf("word is not as expected. Expected: %s, Got: %s", expectedWord, word)
				}
			}
		})
	}
}
