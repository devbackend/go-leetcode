package medium_test

import (
	"github.com/devbackend/go-leetcode/pkg/medium"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLargestNumber(t *testing.T) {
	cases := []struct {
		name     string
		nums     []int
		expected string
	}{
		{
			name:     "example 1",
			nums:     []int{10, 2},
			expected: "210",
		},
		{
			name:     "example 2",
			nums:     []int{3, 30, 34, 5, 9},
			expected: "9534330",
		},
		{
			name:     "example 3",
			nums:     []int{1},
			expected: "1",
		},
		{
			name:     "example 4",
			nums:     []int{10},
			expected: "10",
		},
		{
			name:     "example 5",
			nums:     []int{0, 0},
			expected: "0",
		},
		{
			name:     "example 6",
			nums:     []int{0, 1, 0},
			expected: "100",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := medium.LargestNumber(c.nums)

			assert.Equal(t, c.expected, actual)
		})
	}
}
