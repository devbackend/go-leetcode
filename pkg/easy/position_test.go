package easy_test

import (
	"testing"

	"github.com/devbackend/go-leetcode/pkg/easy"
	"github.com/stretchr/testify/assert"
)

func TestSearchInsert(t *testing.T) {
	cases := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "example 1",
			nums:     []int{1, 3, 5, 6},
			target:   5,
			expected: 2,
		},
		{
			name:     "example 2",
			nums:     []int{1, 3, 5, 6},
			target:   2,
			expected: 1,
		},
		{
			name:     "example 3",
			nums:     []int{1, 3, 5, 6},
			target:   7,
			expected: 4,
		},
		{
			name:     "example 4",
			nums:     []int{1, 3, 5, 6},
			target:   0,
			expected: 0,
		},
		{
			name:     "example 5",
			nums:     []int{1},
			target:   0,
			expected: 0,
		},
		{
			name:     "example 6",
			nums:     []int{1, 2, 3, 4, 5, 10},
			target:   2,
			expected: 1,
		},
	}
	for _, c := range cases {
		if c.name == "" {
			t.Errorf("test case name required!")
			continue
		}

		t.Run(c.name, func(t *testing.T) {
			actual := easy.SearchInsert(c.nums, c.target)
			assert.Equal(t, c.expected, actual)
		})
	}
}
