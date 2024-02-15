package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProblemTopKFrequent(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		expected []int
		nums     []int
		k        int
	}{
		"example 1": {
			expected: []int{1, 2},
			nums:     []int{1, 1, 1, 2, 2, 3},
			k:        2,
		},
		"example 2": {
			expected: []int{1},
			nums:     []int{1},
			k:        1,
		},
		"example 3": {
			expected: []int{1, 2},
			nums:     []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3},
			k:        2,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			require.ElementsMatch(t, tc.expected, topKFrequent(tc.nums, tc.k))
		})
	}
}
