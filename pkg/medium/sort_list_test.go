package medium_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/devbackend/go-leetcode/pkg/medium"
)

func TestSortList(t *testing.T) {
	testCases := map[string]struct {
		head     *medium.ListNode
		expected *medium.ListNode
	}{
		"example 1": {
			head:     buildNodeList([]int{4, 2, 1, 3}),
			expected: buildNodeList([]int{1, 2, 3, 4}),
		},
		"example 2": {
			head:     buildNodeList([]int{-1, 5, 3, 4, 0}),
			expected: buildNodeList([]int{-1, 0, 3, 4, 5}),
		},
		"example 3": {
			head:     buildNodeList([]int{2, 4, 6, 8, 10, 9, 7, 5, 3, 1}),
			expected: buildNodeList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
		},
		"example 4": {
			head:     buildNodeList([]int{3, 1}),
			expected: buildNodeList([]int{1, 3}),
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := medium.SortList(tc.head)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestMergeTwoLists(t *testing.T) {
	testCases := map[string]struct {
		first    *medium.ListNode
		second   *medium.ListNode
		expected *medium.ListNode
	}{
		"only first": {
			first:    buildNodeList([]int{1, 2, 3, 4, 5}),
			expected: buildNodeList([]int{1, 2, 3, 4, 5}),
		},
		"only second": {
			second:   buildNodeList([]int{1, 2, 3, 4, 5}),
			expected: buildNodeList([]int{1, 2, 3, 4, 5}),
		},
		"long first": {
			first:    buildNodeList([]int{1, 2, 3, 7, 8}),
			second:   buildNodeList([]int{4, 5, 6}),
			expected: buildNodeList([]int{1, 2, 3, 4, 5, 6, 7, 8}),
		},
		"long second": {
			first:    buildNodeList([]int{1, 2, 3, 8}),
			second:   buildNodeList([]int{3, 5, 6, 9, 10, 11}),
			expected: buildNodeList([]int{1, 2, 3, 3, 5, 6, 8, 9, 10, 11}),
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := medium.MergeTwoLists(tc.first, tc.second)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
