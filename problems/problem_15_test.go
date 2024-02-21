package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProblemThreeSum(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		nums     []int
		expected [][]int
	}{
		"example 1": {
			nums:     []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		"example 2": {
			nums:     []int{},
			expected: [][]int{},
		},
		"example 3": {
			nums:     []int{0},
			expected: [][]int{},
		},
		"example 4": {
			nums:     []int{0, -1, 1},
			expected: [][]int{{-1, 0, 1}},
		},
		"example 5": {
			nums:     []int{1, -1},
			expected: [][]int{},
		},
		"example 6": {
			nums:     []int{-1, -1, 2, 4, -3, -1, 2, 1},
			expected: [][]int{{-3, -1, 4}, {-3, 1, 2}, {-1, -1, 2}},
		},
		"example 7": {
			nums:     []int{-1, -1, 2, -1, -1, 2},
			expected: [][]int{{-1, -1, 2}},
		},
		"example 8": {
			nums:     []int{0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
		"example 9": {
			nums:     []int{1, 2, -2, -1},
			expected: [][]int{},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			require.ElementsMatch(t, tc.expected, threeSum(tc.nums))
		})
	}
}
