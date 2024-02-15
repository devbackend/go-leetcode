package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		str      string
		expected int
	}{
		"example 1": {
			str:      "abcabcbb",
			expected: 3,
		},
		"example 2": {
			str:      "bbbbb",
			expected: 1,
		},
		"example 3": {
			str:      "pwwkew",
			expected: 3,
		},
		"example 4": {
			str:      "",
			expected: 0,
		},
		"example 5": {
			str:      "aab",
			expected: 2,
		},
		"example 6": {
			str:      "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz 0123456789 abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz",
			expected: 37,
		},
		"example 7": {
			str:      "abc",
			expected: 3,
		},
		"example 8": {
			str:      "abcad",
			expected: 4,
		},
		"example 9": {
			str:      "abcccccca",
			expected: 3,
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			require.Equal(t, tc.expected, lengthOfLongestSubstring(tc.str))
		})
	}
}
