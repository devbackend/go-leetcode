package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProblem_isPalindrome(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		str      string
		expected bool
	}{
		"example 1": {
			str:      "A man, a plan, a canal: Panama",
			expected: true,
		},
		"example 2": {
			str:      "race a car",
			expected: false,
		},
		"example 3": {
			str:      "abba",
			expected: true,
		},
		"example 4": {
			str:      "abcbac",
			expected: false,
		},
		"example 5": {
			str:      "",
			expected: true,
		},
		"example 6": {
			str:      "a",
			expected: true,
		},
		"example 7": {
			str:      "ab",
			expected: false,
		},
		"example 8": {
			str:      "0P",
			expected: false,
		},
		"example 9": {
			str:      "1001",
			expected: true,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			require.Equal(t, tc.expected, isPalindrome(tc.str))
		})
	}
}
