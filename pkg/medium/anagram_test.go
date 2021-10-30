package medium_test

import (
	"testing"

	"github.com/devbackend/go-leetcode/pkg/medium"
	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams(t *testing.T) {
	cases := []struct {
		name     string
		input    []string
		expected [][]string
	}{
		{
			name:     "example 1",
			input:    []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{
			name:     "example 2",
			input:    []string{""},
			expected: [][]string{{""}},
		},
		{
			name:     "example 3",
			input:    []string{"a"},
			expected: [][]string{{"a"}},
		},
		{
			name:     "example 4",
			input:    []string{"abba", "baba", "bbaa", "abcdefghijklmnopqrstuvwxyzzyxwvutsrqponmlkjihgfedcbaabcdefghijklmnopqrstuvwxyzzyxwvutsrqponmlkjihgfedcba"},
			expected: [][]string{{"abba", "baba", "bbaa"}, {"abcdefghijklmnopqrstuvwxyzzyxwvutsrqponmlkjihgfedcbaabcdefghijklmnopqrstuvwxyzzyxwvutsrqponmlkjihgfedcba"}},
		},
	}

	for _, c := range cases {
		if c.name == "" {
			t.Errorf("test case name required!")
			continue
		}

		t.Run(c.name, func(t *testing.T) {
			actual := medium.GroupAnagrams(c.input)

			assert.Equal(t, c.expected, actual)
		})
	}
}
