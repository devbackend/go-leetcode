package easy_test

import (
	"testing"

	"github.com/devbackend/go-leetcode/pkg/easy"
	"github.com/stretchr/testify/assert"
)

func TestIsPalindromeNumber(t *testing.T) {
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
		if c.name == "" {
			t.Errorf("test case name required!")
			continue
		}

		t.Run(c.name, func(t *testing.T) {
			actual := easy.IsPalindromeNumber(c.num)
			assert.Equal(t, c.expected, actual)
		})
	}
}

func TestIsPalindromeString(t *testing.T) {
	cases := []struct {
		name     string
		str      string
		expected bool
	}{
		{
			name:     "example 1",
			str:      "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			name:     "example 2",
			str:      "race a car",
			expected: false,
		},
		{
			name:     "example 3",
			str:      "abba",
			expected: true,
		},
		{
			name:     "example 4",
			str:      "abcbac",
			expected: false,
		},
		{
			name:     "example 5",
			str:      "",
			expected: true,
		},
		{
			name:     "example 6",
			str:      "a",
			expected: true,
		},
		{
			name:     "example 7",
			str:      "ab",
			expected: false,
		},
		{
			name:     "example 8",
			str:      "0P",
			expected: false,
		},
		{
			name:     "example 9",
			str:      "1001",
			expected: true,
		},
	}
	for _, c := range cases {
		if c.name == "" {
			t.Errorf("test case name required!")
			continue
		}

		t.Run(c.name, func(t *testing.T) {
			actual := easy.IsPalindromeString(c.str)
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
