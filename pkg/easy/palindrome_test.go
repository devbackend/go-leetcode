package easy_test

import (
	"testing"

	"github.com/devbackend/go-leetcode/pkg/easy"
	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		name     string
		num      int
		expected bool
	}{
		{
			name:     "example 1",
			num:      121,
			expected: true,
		},
		{
			name:     "example 2",
			num:      -121,
			expected: false,
		},
		{
			name:     "example 3",
			num:      10,
			expected: false,
		},
		{
			name:     "example 4",
			num:      -101,
			expected: false,
		},
	}
	for _, c := range cases {
		c := c // scopelint mute

		if c.name == "" {
			t.Errorf("test case name required!")
			continue
		}

		t.Run(c.name, func(t *testing.T) {
			actual := easy.IsPalindrome(c.num)
			assert.Equal(t, c.expected, actual)
		})
	}
}

func TestIsPalindromeLinkedList(t *testing.T) {
	cases := []struct {
		name     string
		head     *easy.ListNode
		expected bool
	}{
		{
			name:     "example 1",
			expected: true,
			head: &easy.ListNode{
				Val: 1,
				Next: &easy.ListNode{
					Val: 2,
					Next: &easy.ListNode{
						Val: 2,
						Next: &easy.ListNode{
							Val: 1,
						},
					},
				},
			},
		},
		{
			name:     "example 3",
			expected: true,
			head: &easy.ListNode{
				Val: 1,
			},
		},
		{
			name:     "example 4",
			expected: true,
		},
	}
	for _, c := range cases {
		c := c // scopelint mute

		if c.name == "" {
			t.Errorf("test case name required!")
			continue
		}

		t.Run(c.name, func(t *testing.T) {
			actual := easy.IsPalindromeLinkedList(c.head)
			assert.Equal(t, c.expected, actual)
		})
	}
}
