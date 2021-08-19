package easy_test

import (
	"testing"

	"github.com/devbackend/go-leetcode/pkg/easy"
	"github.com/stretchr/testify/assert"
)

func TestHasPathSum(t *testing.T) {
	cases := []struct {
		name     string
		tree     *easy.TreeNode
		target   int
		expected bool
	}{
		{
			name:     "example 1",
			target:   22,
			expected: true,
			tree: &easy.TreeNode{
				Val: 5,
				Left: &easy.TreeNode{
					Val: 4,
					Left: &easy.TreeNode{
						Val: 11,
						Left: &easy.TreeNode{
							Val: 8,
						},
						Right: &easy.TreeNode{
							Val: 2,
						},
					},
				},
				Right: &easy.TreeNode{
					Val: 8,
					Left: &easy.TreeNode{
						Val: 13,
					},
					Right: &easy.TreeNode{
						Val: 8,
						Right: &easy.TreeNode{
							Val: 4,
							Right: &easy.TreeNode{
								Val: 1,
							},
						},
					},
				},
			},
		},
		{
			name:     "example 2",
			target:   5,
			expected: false,
			tree: &easy.TreeNode{
				Val: 1,
				Right: &easy.TreeNode{
					Val: 2,
				},
				Left: &easy.TreeNode{
					Val: 2,
				},
			},
		},
		{
			name:     "example 3",
			target:   1,
			expected: false,
			tree: &easy.TreeNode{
				Val: 1,
				Right: &easy.TreeNode{
					Val: 2,
				},
			},
		},
		{
			name:     "example 4",
			target:   1,
			expected: true,
			tree: &easy.TreeNode{
				Val: 1,
			},
		},
		{
			name:     "example 5",
			target:   -5,
			expected: true,
			tree: &easy.TreeNode{
				Val:   -2,
				Right: &easy.TreeNode{Val: -3},
			},
		},
		{
			name:     "example 6",
			target:   -1,
			expected: true,
			tree: &easy.TreeNode{
				Val: 1,
				Right: &easy.TreeNode{
					Val: -2,
					Left: &easy.TreeNode{
						Val: 1,
						Left: &easy.TreeNode{
							Val: -1,
						},
					},
					Right: &easy.TreeNode{
						Val: 3,
					},
				},
				Left: &easy.TreeNode{
					Val: -3,
					Left: &easy.TreeNode{
						Val: -2,
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
			actual := easy.HasPathSum(c.tree, c.target)
			assert.Equal(t, c.expected, actual)
		})
	}
}
