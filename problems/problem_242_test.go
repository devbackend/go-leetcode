package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProblemIsAnagram(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		expected bool
		first    string
		second   string
	}{
		"example 1": {
			expected: true,
			first:    "anagram",
			second:   "nagaram",
		},
		"example 2": {
			expected: false,
			first:    "rat",
			second:   "car",
		},
		"example 3": {
			first:    "abc",
			second:   "cab",
			expected: true,
		},
		"example 4": {
			first:    "thanos",
			second:   "sonata",
			expected: false,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			require.Equal(t, tc.expected, isAnagram(tc.first, tc.second))
		})
	}
}
