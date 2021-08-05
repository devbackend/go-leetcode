package medium_test

import (
	"github.com/devbackend/go-leetcode/pkg/medium"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	cases := []struct {
		name     string
		first    *medium.ListNode
		second   *medium.ListNode
		expected *medium.ListNode
	}{
		{
			name:     "example 1",
			first:    buildNodeList([]int{2, 4, 3}),
			second:   buildNodeList([]int{5, 6, 4}),
			expected: buildNodeList([]int{7, 0, 8}),
		},
		{
			name:     "example 2",
			first:    buildNodeList([]int{0}),
			second:   buildNodeList([]int{0}),
			expected: buildNodeList([]int{0}),
		},
		{
			name:     "example 3",
			first:    buildNodeList([]int{9, 9, 9, 9, 9, 9, 9}),
			second:   buildNodeList([]int{9, 9, 9, 9}),
			expected: buildNodeList([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
		{
			name:     "example 4",
			first:    buildNodeList([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}),
			second:   buildNodeList([]int{5, 6, 4}),
			expected: buildNodeList([]int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}),
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := medium.AddTwoNumbers(c.first, c.second)

			assert.Equal(t, c.expected, actual)
		})
	}
}

func buildNodeList(nums []int) *medium.ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &medium.ListNode{Val: nums[0]}

	curr := head
	for i := 1; i < len(nums); i++ {
		curr.Next = &medium.ListNode{Val: nums[i]}

		curr = curr.Next
	}

	return head
}
