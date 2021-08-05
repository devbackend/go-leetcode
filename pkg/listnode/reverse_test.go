package listnode_test

import (
	"github.com/devbackend/go-leetcode/pkg/listnode"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		name     string
		listNode *listnode.ListNode
		expected *listnode.ListNode
	}{
		{
			name: "empty",
		},
		{
			name:     "one node",
			listNode: &listnode.ListNode{Val: 1},
			expected: &listnode.ListNode{Val: 1},
		},
		{
			name: "two nodes",
			listNode: &listnode.ListNode{
				Val:  1,
				Next: &listnode.ListNode{Val: 2},
			},
			expected: &listnode.ListNode{
				Val: 2,
				Next: &listnode.ListNode{
					Val: 1,
				},
			},
		},
		{
			name: "three nodes",
			listNode: &listnode.ListNode{
				Val: 1,
				Next: &listnode.ListNode{
					Val: 2,
					Next: &listnode.ListNode{
						Val: 3,
					},
				},
			},
			expected: &listnode.ListNode{
				Val: 3,
				Next: &listnode.ListNode{
					Val: 2,
					Next: &listnode.ListNode{
						Val: 1,
					},
				},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := listnode.Reverse(c.listNode)
			assert.Equal(t, c.expected, actual)
		})
	}
}
