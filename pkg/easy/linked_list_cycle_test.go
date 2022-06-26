package easy_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/devbackend/go-leetcode/pkg/easy"
)

func TestHasCycle(t *testing.T) {
	testCases := map[string]struct {
		head     *easy.ListNode
		expected bool
	}{
		"empty": {
			head:     nil,
			expected: false,
		},
		"only head": {
			head:     &easy.ListNode{},
			expected: false,
		},
		"no cycled": {
			head: &easy.ListNode{
				Val: 1,
				Next: &easy.ListNode{
					Val: 2,
				},
			},
		},
		"has cycle": {
			expected: true,
			head: func() *easy.ListNode {
				tail := &easy.ListNode{
					Val: 1000,
				}

				head := &easy.ListNode{
					Val: 1,
					Next: &easy.ListNode{
						Val: 2,
						Next: &easy.ListNode{
							Val: 3,
							Next: &easy.ListNode{
								Val: 4,
								Next: &easy.ListNode{
									Val: 5,
									Next: &easy.ListNode{
										Val: 6,
										Next: &easy.ListNode{
											Val:  7,
											Next: tail,
										},
									},
								},
							},
						},
					},
				}

				tail.Next = head.Next.Next.Next

				return head
			}(),
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := easy.HasCycle(tc.head)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
