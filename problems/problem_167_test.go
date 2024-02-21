package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProblem_twoSum(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		nums     []int
		target   int
		expected []int
	}{
		"example 1": {
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{1, 2},
		},
		"example 2": {
			nums:     []int{2, 3, 4},
			target:   6,
			expected: []int{1, 3},
		},
		"example 3": {
			nums:     []int{-1, 0},
			target:   -1,
			expected: []int{1, 2},
		},
		"example 4": {
			nums:     []int{3, 24, 50, 79, 88, 150, 345},
			target:   200,
			expected: []int{3, 6},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			require.Equal(t, tc.expected, twoSumSorted(tc.nums, tc.target))
		})
	}
}
