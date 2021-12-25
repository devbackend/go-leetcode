package easy_test

import (
	"testing"

	"github.com/devbackend/go-leetcode/pkg/easy"
	"github.com/stretchr/testify/assert"
)

func TestMaxSubArray(t *testing.T) {
	testCases := map[string]struct {
		nums     []int
		expected int
	}{
		"example 1": {nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, expected: 6},
		"example 2": {nums: []int{1}, expected: 1},
		"example 3": {nums: []int{5, 4, -1, 7, 8}, expected: 23},
		"example 4": {nums: []int{-1}, expected: -1},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := easy.MaxSubArray(tc.nums)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
