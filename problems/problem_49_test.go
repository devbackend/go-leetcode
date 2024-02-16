package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProblemGroupAnagrams(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		strs     []string
		expected [][]string
	}{
		"example 1": {
			strs:     []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		"example 2": {
			strs:     []string{""},
			expected: [][]string{{""}},
		},
		"example 3": {
			strs:     []string{"a"},
			expected: [][]string{{"a"}},
		},
		"example 4": {
			strs:     []string{"abba", "baba", "bbaa", "abcdefghijklmnopqrstuvwxyzzyxwvutsrqponmlkjihgfedcbaabcdefghijklmnopqrstuvwxyzzyxwvutsrqponmlkjihgfedcba"},
			expected: [][]string{{"abba", "baba", "bbaa"}, {"abcdefghijklmnopqrstuvwxyzzyxwvutsrqponmlkjihgfedcbaabcdefghijklmnopqrstuvwxyzzyxwvutsrqponmlkjihgfedcba"}},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			require.ElementsMatch(t, tc.expected, groupAnagrams(tc.strs))
		})
	}
}
