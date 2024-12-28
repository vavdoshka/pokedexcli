package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct{
		input string
		expected []string
	}{
		{
		input: " what a Nice day ",
		expected: []string{"what", "a", "nice", "day"},
		},
		{
			input: "SOMETHINGLARGE",
			expected: []string{"somethinglarge"},
		},
		{
			input: "",
			expected: []string{},
		},
		{
			input: "what\ta\tnice\nday",
			expected: []string{"what", "a", "nice", "day"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("expected length: %d - (%v) does not match actual: %d - (%v)", len(c.expected), c.expected, len(actual), actual)
			t.Fail()
			continue
		}

		for i := range c.expected {
			expectedWord := c.expected[i]
			word := actual[i]

			if word != expectedWord {
				t.Errorf("expected: '%s' does not match actual '%s'", expectedWord, word)
				t.Fail()
			}
		}
	}
}