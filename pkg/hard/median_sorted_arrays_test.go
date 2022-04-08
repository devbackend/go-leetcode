package hard_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/devbackend/go-leetcode/pkg/hard"
)

func TestFindMedianSortedArrays(t *testing.T) {
	testCases := map[string]struct {
		nums1    []int
		nums2    []int
		expected float64
	}{
		"example 1": {
			nums1:    []int{1, 3},
			nums2:    []int{2},
			expected: 2,
		},
		"example 2": {
			nums1:    []int{1, 2},
			nums2:    []int{3, 4},
			expected: 2.5,
		},
		"example 3": {
			nums1:    []int{3, 4},
			nums2:    []int{5, 6},
			expected: 4.5,
		},
		"example 4": {
			nums1:    []int{8},
			nums2:    []int{7, 9},
			expected: 8,
		},
		"example 5": {
			nums1:    []int{1, 3, 5},
			nums2:    []int{2, 4},
			expected: 3,
		},
		"example 6": {
			nums1:    []int{1, 3, 5},
			nums2:    []int{2, 4, 6},
			expected: 3.5,
		},
		"example 7": {
			nums1:    []int{1, 2, 3},
			nums2:    []int{1, 2, 3},
			expected: 2,
		},
		"example 8": {
			nums1:    []int{1, 2, 3},
			nums2:    []int{2, 3},
			expected: 2,
		},
		"example 9": {
			nums1: func() []int {
				res := make([]int, 999)

				for i := range res {
					res[i] = i + 1
				}

				return res
			}(),
			nums2: func() []int {
				res := make([]int, 999)

				for i := range res {
					res[i] = i + 1000
				}

				return res
			}(),
			expected: 999.5,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := hard.FindMedianSortedArrays(tc.nums1, tc.nums2)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
