package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProblemProductExceptSelf(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		expected []int
		nums     []int
	}{
		"example 1": {
			nums:     []int{1, 2, 3, 4},
			expected: []int{24, 12, 8, 6},
		},
		"example 2": {
			nums:     []int{-1, 1, 0, -3, 3},
			expected: []int{0, 0, 9, 0, 0},
		},
		"example 3": {
			nums:     []int{5, 10, 20, 50},
			expected: []int{10000, 5000, 2500, 1000},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			require.Equal(t, tc.expected, productExceptSelf(tc.nums))
		})
	}
}
