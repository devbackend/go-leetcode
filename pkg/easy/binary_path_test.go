package easy_test

import (
	"sort"
	"testing"

	"github.com/devbackend/go-leetcode/pkg/easy"
	"github.com/stretchr/testify/assert"
)

func TestBinaryTreePaths(t *testing.T) {
	cases := []struct {
		name     string
		root     *easy.TreeNode
		expected []string
	}{
		{
			name: "example 1",
			root: &easy.TreeNode{
				Val: 1,
				Left: &easy.TreeNode{
					Val:   2,
					Right: &easy.TreeNode{Val: 5},
				},
				Right: &easy.TreeNode{Val: 3},
			},
			expected: []string{"1->2->5", "1->3"},
		},
		{
			name:     "example 2",
			root:     &easy.TreeNode{Val: 1},
			expected: []string{"1"},
		},
		{
			name: "example 3",
			root: &easy.TreeNode{
				Val:   1,
				Left:  &easy.TreeNode{Val: 2},
				Right: &easy.TreeNode{Val: 3},
			},
			expected: []string{"1->2", "1->3"},
		},
		{
			name: "example 4",
			root: &easy.TreeNode{
				Val: 1,
				Left: &easy.TreeNode{
					Val:   2,
					Left:  &easy.TreeNode{Val: 4},
					Right: &easy.TreeNode{Val: 5},
				},
				Right: &easy.TreeNode{
					Val:   3,
					Left:  &easy.TreeNode{Val: 6},
					Right: &easy.TreeNode{Val: 7},
				},
			},
			expected: []string{"1->2->4", "1->2->5", "1->3->6", "1->3->7"},
		},
	}
	for _, c := range cases {
		if c.name == "" {
			t.Errorf("test case name required!")
			continue
		}

		t.Run(c.name, func(t *testing.T) {
			actual := easy.BinaryTreePaths(c.root)

			sort.Strings(actual)
			sort.Strings(c.expected)

			assert.Equal(t, c.expected, actual)
		})
	}
}
