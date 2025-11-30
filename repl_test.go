package main

import (
	"reflect"
	"testing"
)

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
			input:    "hey how are you   ",
			expected: []string{"hey", "how", "are", "you"},
		},
		{
			input:    "  i'm fine    thank    you",
			expected: []string{"i'm", "fine", "thank", "you"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if !reflect.DeepEqual(len(actual), len(c.expected)) {
			t.Errorf("Length of actual %d doesn't match with expected %d", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expected := c.expected[i]
			if !reflect.DeepEqual(word, expected) {
				t.Errorf("Words at position %d doesn't match\n actual: %s\n expected: %s", i, word, expected)
			}
		}
	}

}
