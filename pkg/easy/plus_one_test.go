package easy_test

import (
	"testing"

	"github.com/devbackend/go-leetcode/pkg/easy"
	"github.com/stretchr/testify/assert"
)

func TestPlusOne(t *testing.T) {
	cases := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "example 1",
			nums:     []int{1, 2, 3},
			expected: []int{1, 2, 4},
		},
		{
			name:     "example 2",
			nums:     []int{4, 3, 2, 1},
			expected: []int{4, 3, 2, 2},
		},
		{
			name:     "example 3",
			nums:     []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
			expected: []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	for _, c := range cases {
		if c.name == "" {
			t.Errorf("test case name required!")
			continue
		}

		t.Run(c.name, func(t *testing.T) {
			actual := easy.PlusOne(c.nums)
			assert.Equal(t, c.expected, actual)
		})
	}
}
