package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProblemMaxArea(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		heights  []int
		expected int
	}{
		"example 1": {
			heights:  []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		"example 2": {
			heights:  []int{4, 3, 2, 1, 4},
			expected: 16,
		},
		"example 3": {
			heights:  []int{1, 2, 1},
			expected: 2,
		},
		"example 4": {
			expected: 233767415,
			heights: func() []int {
				res := make([]int, 96700)

				for i := 0; i < 96700; i++ {
					res[i] = i / 10
				}

				return res
			}(),
		},
		"example 5": {
			heights:  []int{2, 2},
			expected: 2,
		},
		"example 6": {
			heights:  []int{2, 0, 2, 1},
			expected: 4,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			require.Equal(t, tc.expected, maxArea(tc.heights))
		})
	}
}
