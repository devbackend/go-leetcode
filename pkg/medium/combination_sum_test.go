package medium_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/devbackend/go-leetcode/pkg/medium"
)

func TestCombinationSum(t *testing.T) {
	testCases := map[string]struct {
		candidates []int
		target     int
		expected   [][]int
	}{
		"example 1": {
			candidates: []int{2, 3, 6, 7},
			target:     7,
			expected: [][]int{
				{2, 2, 3},
				{7},
			},
		},
		"example 2": {
			candidates: []int{2},
			target:     1,
			expected:   [][]int{},
		},
		"example 3": {
			candidates: []int{1, 2, 3, 4, 5},
			target:     4,
			expected: [][]int{
				{1, 1, 1, 1},
				{1, 1, 2},
				{1, 3},
				{2, 2},
				{4},
			},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := medium.CombinationSum(tc.candidates, tc.target)
			assert.ElementsMatch(t, tc.expected, actual)
		})
	}
}
