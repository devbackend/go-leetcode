package easy_test

import (
	"testing"

	"github.com/devbackend/go-leetcode/pkg/easy"
	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	cases := []struct {
		name     string
		nums1    []int
		m        int
		nums2    []int
		n        int
		expected []int
	}{
		{
			name:     "example 1",
			nums1:    []int{1, 2, 3, 0, 0, 0},
			m:        3,
			nums2:    []int{2, 5, 6},
			n:        3,
			expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name:     "example 2",
			nums1:    []int{1},
			m:        1,
			nums2:    []int{},
			n:        0,
			expected: []int{1},
		},
		{
			name:     "example 3",
			nums1:    []int{0},
			m:        0,
			nums2:    []int{1},
			n:        1,
			expected: []int{1},
		},
		{
			name:     "example 4",
			nums1:    []int{-1, 0, 1, 0, 0},
			m:        3,
			nums2:    []int{2, 3},
			n:        2,
			expected: []int{-1, 0, 1, 2, 3},
		},
		{
			name:     "example 5",
			nums1:    []int{4, 5, 6, 0, 0, 0},
			m:        3,
			nums2:    []int{1, 2, 3},
			n:        3,
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "example 6",
			nums1:    []int{1, 2, 4, 5, 6, 0},
			m:        5,
			nums2:    []int{3},
			n:        1,
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "example 7",
			nums1:    []int{0, 0, 3, 0, 0, 0, 0, 0, 0},
			m:        3,
			nums2:    []int{-1, 1, 1, 1, 2, 3},
			n:        6,
			expected: []int{-1, 0, 0, 1, 1, 1, 2, 3, 3},
		},
		{
			name:     "example 8",
			nums1:    []int{-1, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0},
			m:        5,
			nums2:    []int{-1, -1, 0, 0, 1, 2},
			n:        6,
			expected: []int{-1, -1, -1, 0, 0, 0, 0, 0, 1, 2, 3},
		},
	}
	for _, c := range cases {
		if c.name == "" {
			t.Errorf("test case name required!")
			continue
		}

		t.Run(c.name, func(t *testing.T) {
			easy.Merge(c.nums1, c.m, c.nums2, c.n)
			assert.Equal(t, c.expected, c.nums1)
		})
	}
}
