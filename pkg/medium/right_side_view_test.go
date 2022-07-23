package medium_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/devbackend/go-leetcode/pkg/medium"
)

func TestRightSideView(t *testing.T) {
	testCases := map[string]struct {
		root     *medium.TreeNode
		expected []int
	}{
		"example 1": {
			root: &medium.TreeNode{
				Val: 1,
				Left: &medium.TreeNode{
					Val: 2,
					Right: &medium.TreeNode{
						Val: 5,
					},
				},
				Right: &medium.TreeNode{
					Val: 3,
					Right: &medium.TreeNode{
						Val: 4,
					},
				},
			},
			expected: []int{1, 3, 4},
		},
		"example 2": {
			root: &medium.TreeNode{
				Val: 1,
				Left: &medium.TreeNode{
					Val: 3,
				},
			},
			expected: []int{1, 3},
		},
		"example 3": {
			root:     nil,
			expected: []int(nil),
		},
		"example 4": {
			root:     &medium.TreeNode{Val: 1},
			expected: []int{1},
		},
		"example 5": {
			root: &medium.TreeNode{
				Val: 1,
				Left: &medium.TreeNode{
					Val: 2,
					Left: &medium.TreeNode{
						Val: 4,
						Left: &medium.TreeNode{
							Val: 6,
						},
						Right: &medium.TreeNode{
							Val: 7,
						},
					},
					Right: &medium.TreeNode{
						Val: 5,
					},
				},
				Right: &medium.TreeNode{
					Val: 3,
				},
			},
			expected: []int{1, 3, 5, 7},
		},
		"example 6": {
			root: &medium.TreeNode{
				Val: 1,
				Left: &medium.TreeNode{
					Val: 2,
					Left: &medium.TreeNode{
						Val: 4,
						Left: &medium.TreeNode{
							Val: 7,
						},
					},
					Right: &medium.TreeNode{
						Val: 5,
						Left: &medium.TreeNode{
							Val: 6,
							Left: &medium.TreeNode{
								Val: 8,
							},
						},
					},
				},
				Right: &medium.TreeNode{
					Val: 3,
					Left: &medium.TreeNode{
						Val: 9,
					},
				},
			},
			expected: []int{1, 3, 9, 6, 8},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := medium.RightSideView(tc.root)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
