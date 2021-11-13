package medium_test

import (
	"testing"

	"github.com/devbackend/go-leetcode/pkg/medium"
	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T) {
	cases := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name:     "example 1",
			nums:     []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name:     "example 2",
			nums:     []int{},
			expected: [][]int{},
		},
		{
			name:     "example 3",
			nums:     []int{0},
			expected: [][]int{},
		},
		{
			name:     "example 4",
			nums:     []int{0, -1, 1},
			expected: [][]int{{-1, 0, 1}},
		},
		{
			name:     "example 5",
			nums:     []int{1, -1},
			expected: [][]int{},
		},
		{
			name:     "example 6",
			nums:     []int{-1, -1, 2, 4, -3, -1, 2, 1},
			expected: [][]int{{-3, -1, 4}, {-3, 1, 2}, {-1, -1, 2}},
		},
		{
			name:     "example 7",
			nums:     []int{-1, -1, 2, -1, -1, 2},
			expected: [][]int{{-1, -1, 2}},
		},
		{
			name:     "example 8",
			nums:     []int{0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
		{
			name:     "example 9",
			nums:     []int{1, 2, -2, -1},
			expected: [][]int{},
		},
	}
	for _, c := range cases {
		if c.name == "" {
			t.Errorf("test case name required!")
			continue
		}

		t.Run(c.name, func(t *testing.T) {
			actual := medium.ThreeSum(c.nums)
			assert.ElementsMatch(t, c.expected, actual)
		})
	}
}
