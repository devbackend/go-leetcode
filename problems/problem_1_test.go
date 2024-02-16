package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	cases := map[string]struct {
		nums     []int
		target   int
		expected []int
	}{
		"example 1": {
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		"example 2": {
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		"example 3": {
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
		"example 4": {
			nums:     []int{-3, 0, 3, -7},
			target:   -10,
			expected: []int{0, 3},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			actual := twoSum(tc.nums, tc.target)

			assert.Equal(t, tc.expected, actual)
		})
	}
}
