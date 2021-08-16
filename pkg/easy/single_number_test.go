package easy_test

import (
	"github.com/devbackend/go-leetcode/pkg/easy"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	cases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "example 1",
			nums:     []int{2, 2, 1},
			expected: 1,
		},
		{
			name:     "example 2",
			nums:     []int{4, 1, 2, 1, 2},
			expected: 4,
		},
		{
			name:     "example 3",
			nums:     []int{1},
			expected: 1,
		},
		{
			name:     "example 4",
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 11},
			expected: 11,
		},
	}
	for _, c := range cases {
		if c.name == "" {
			t.Errorf("test case name required!")
			continue
		}

		t.Run(c.name, func(t *testing.T) {
			actual := easy.SingleNumber(c.nums)
			assert.Equal(t, c.expected, actual)
		})
	}
}
