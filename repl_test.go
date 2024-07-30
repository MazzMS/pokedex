package main

import "testing"

func TestCleanInput(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{
			input: "testing THIS func", expected: []string{
				"testing",
				"this",
				"func",
			},
		},
	}

	for _, test := range tests {
		actual := clearInput(test.input)
		if len(actual) != len(test.expected) {
			t.Errorf("The lengths are not equal: %d vs %d",
				len(actual),
				len(test.expected),
			)
			continue
		}
		for i := range actual {
			if actual[i] != test.expected[i] {
				t.Errorf("%q is not %q",
					actual[i],
					test.expected[i],
				)
			}
		}
	}
}
