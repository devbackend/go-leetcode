package listnode_test

import (
	"testing"

	"github.com/devbackend/go-leetcode/pkg/listnode"
	"github.com/stretchr/testify/assert"
)

func TestMergeTwoLists(t *testing.T) {
	cases := []struct {
		name     string
		first    *listnode.ListNode
		second   *listnode.ListNode
		expected *listnode.ListNode
	}{
		{
			name:     "example 1",
			first:    listnode.BuildFromSlice([]int{1, 2, 4}),
			second:   listnode.BuildFromSlice([]int{1, 3, 4}),
			expected: listnode.BuildFromSlice([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			name:     "example 2",
			first:    listnode.BuildFromSlice([]int{}),
			second:   listnode.BuildFromSlice([]int{}),
			expected: listnode.BuildFromSlice([]int{}),
		},
		{
			name:     "example 3",
			first:    listnode.BuildFromSlice([]int{}),
			second:   listnode.BuildFromSlice([]int{0}),
			expected: listnode.BuildFromSlice([]int{0}),
		},
		{
			name:     "example 4",
			first:    listnode.BuildFromSlice([]int{1}),
			second:   listnode.BuildFromSlice([]int{2}),
			expected: listnode.BuildFromSlice([]int{1, 2}),
		},
		{
			name:     "example 5",
			first:    listnode.BuildFromSlice([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}),
			second:   listnode.BuildFromSlice([]int{2, 3, 4}),
			expected: listnode.BuildFromSlice([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 3, 4}),
		},
	}
	for _, c := range cases {
		if c.name == "" {
			t.Errorf("test case name required!")
			continue
		}

		t.Run(c.name, func(t *testing.T) {
			actual := listnode.MergeTwoLists(c.first, c.second)
			assert.Equal(t, c.expected, actual)
		})
	}
}
