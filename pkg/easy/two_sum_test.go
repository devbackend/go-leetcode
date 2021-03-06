package easy_test

import (
	"testing"

	"github.com/devbackend/go-leetcode/pkg/easy"
	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	cases := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "example 1",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "example 2",
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			name:     "example 3",
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
		{
			name:     "example 4",
			nums:     []int{-3, 0, 3, -7},
			target:   -10,
			expected: []int{0, 3},
		},
	}

	for _, c := range cases {
		if c.name == "" {
			t.Errorf("test case name required!")
			continue
		}

		t.Run(c.name, func(t *testing.T) {
			actual := easy.TwoSum(c.nums, c.target)

			assert.Equal(t, c.expected, actual)
		})
	}
}

func TestTwoSumSorted(t *testing.T) {
	cases := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "example 1",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{1, 2},
		},
		{
			name:     "example 2",
			nums:     []int{2, 3, 4},
			target:   6,
			expected: []int{1, 3},
		},
		{
			name:     "example 3",
			nums:     []int{-1, 0},
			target:   -1,
			expected: []int{1, 2},
		},
	}
	for _, c := range cases {
		if c.name == "" {
			t.Errorf("test case name required!")
			continue
		}

		t.Run(c.name, func(t *testing.T) {
			actual := easy.TwoSumSorted(c.nums, c.target)
			assert.Equal(t, c.expected, actual)
		})
	}
}
