package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContainsDuplicate(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		expected bool
		nums     []int
	}{
		"example 1": {
			expected: true,
			nums:     []int{1, 2, 3, 1},
		},
		"example 2": {
			expected: false,
			nums:     []int{1, 2, 3, 4},
		},
		"example 3": {
			expected: true,
			nums:     []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			require.Equal(t, tc.expected, containsDuplicate(tc.nums))
		})
	}
}
