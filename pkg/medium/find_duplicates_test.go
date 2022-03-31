package medium_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/devbackend/go-leetcode/pkg/medium"
)

func TestFindDuplicate(t *testing.T) {
	testCases := map[string]struct {
		nums     []int
		expected int
	}{
		"example 1": {
			nums:     []int{1, 3, 4, 2, 2},
			expected: 2,
		},
		"example 2": {
			nums:     []int{3, 1, 3, 4, 2},
			expected: 3,
		},
		"example 3": {
			nums:     []int{17, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17},
			expected: 17,
		},
		"example 4": {
			nums:     []int{2, 2, 2, 2, 2},
			expected: 2,
		},
		"example 5": {
			nums:     []int{2, 5, 9, 6, 9, 3, 8, 9, 7, 1},
			expected: 9,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := medium.FindDuplicate(tc.nums)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
