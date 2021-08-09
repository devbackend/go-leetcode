package listnode_test

import (
	"testing"

	"github.com/devbackend/go-leetcode/pkg/listnode"
	"github.com/stretchr/testify/assert"
)

func TestBuildFromSlice(t *testing.T) {
	cases := []struct {
		name     string
		slice    []int
		expected *listnode.ListNode
	}{
		{
			name:     "one element",
			slice:    []int{1},
			expected: &listnode.ListNode{Val: 1},
		},
		{
			name:  "empty",
			slice: []int{},
		},
		{
			name:  "five elements",
			slice: []int{1, 2, 3, 4, 5},
			expected: &listnode.ListNode{
				Val: 1,
				Next: &listnode.ListNode{
					Val: 2,
					Next: &listnode.ListNode{
						Val: 3,
						Next: &listnode.ListNode{
							Val: 4,
							Next: &listnode.ListNode{
								Val: 5,
							},
						},
					},
				},
			},
		},
	}
	for _, c := range cases {
		if c.name == "" {
			t.Errorf("test case name required!")
			continue
		}

		t.Run(c.name, func(t *testing.T) {
			actual := listnode.BuildFromSlice(c.slice)
			assert.Equal(t, c.expected, actual)
		})
	}
}
