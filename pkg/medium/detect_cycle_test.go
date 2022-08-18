package medium_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/devbackend/go-leetcode/pkg/medium"
)

func TestDetectCycle(t *testing.T) {
	type testCase struct {
		head     *medium.ListNode
		expected *medium.ListNode
	}

	testCases := map[string]testCase{
		"example 1": func() testCase {
			listNode := buildNodeList([]int{3, 2, 0, -4})

			listNode.Next.Next.Next.Next = listNode.Next

			return testCase{
				head:     listNode,
				expected: listNode.Next,
			}
		}(),
		"example 2": func() testCase {
			listNode := buildNodeList([]int{1, 2})

			listNode.Next.Next = listNode

			return testCase{
				head:     listNode,
				expected: listNode,
			}
		}(),
		"example 3": func() testCase {
			listNode := buildNodeList([]int{1})

			return testCase{
				head:     listNode,
				expected: nil,
			}
		}(),
		"example 4": func() testCase {
			listNode := buildNodeList([]int{1, 2, 3, 4, 5})

			return testCase{
				head:     listNode,
				expected: nil,
			}
		}(),
		"example 5": func() testCase {
			listNode := buildNodeList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})

			listNode.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = listNode.Next.Next.Next.Next.Next.Next.Next

			return testCase{
				head:     listNode,
				expected: listNode.Next.Next.Next.Next.Next.Next.Next,
			}
		}(),
		"example 6": func() testCase {
			listNode := buildNodeList([]int{1, 2})

			return testCase{
				head:     listNode,
				expected: nil,
			}
		}(),
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := medium.DetectCycle(tc.head)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
