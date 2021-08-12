package medium_test

import (
	"testing"

	"github.com/devbackend/go-leetcode/pkg/medium"
	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	cases := []struct {
		name     string
		str      string
		expected int
	}{
		{
			name:     "example 1",
			str:      "abcabcbb",
			expected: 3,
		},
		{
			name:     "example 2",
			str:      "bbbbb",
			expected: 1,
		},
		{
			name:     "example 3",
			str:      "pwwkew",
			expected: 3,
		},
		{
			name:     "example 4",
			str:      "",
			expected: 0,
		},
		{
			name:     "example 5",
			str:      "aab",
			expected: 2,
		},
		{
			name:     "example 6",
			str:      "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz 0123456789 abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz",
			expected: 37,
		},
		{
			name:     "example 7",
			str:      "abc",
			expected: 3,
		},
		{
			name:     "example 8",
			str:      "abcad",
			expected: 4,
		},
		{
			name:     "example 9",
			str:      "abcccccca",
			expected: 3,
		},
	}
	for _, c := range cases {
		if c.name == "" {
			t.Errorf("test case name required!")
			continue
		}

		t.Run(c.name, func(t *testing.T) {
			actual := medium.LengthOfLongestSubstring(c.str)
			assert.Equal(t, c.expected, actual)
		})
	}
}
