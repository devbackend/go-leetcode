package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProblemLongestConsecutive(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		nums     []int
		expected int
	}{
		"example 1": {
			nums:     []int{100, 4, 200, 1, 3, 2},
			expected: 4,
		},
		"example 2": {
			nums:     []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			expected: 9,
		},
		"example 3": {
			nums:     []int{3, 3, 3, 1, 2, 4},
			expected: 4,
		},
		"example 4": {
			nums:     []int{-2, 101, 3, -1, 0, -3, 2, 1, 7},
			expected: 7,
		},
		"example 5": {
			nums:     []int{10, 20, 30},
			expected: 1,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			require.Equal(t, tc.expected, longestConsecutive(tc.nums))
		})
	}
}
