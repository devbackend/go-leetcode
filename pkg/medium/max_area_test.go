package medium_test

import (
	"testing"

	"github.com/devbackend/go-leetcode/pkg/medium"
	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {
	cases := []struct {
		name     string
		heights  []int
		expected int
	}{
		{
			name:     "example 1",
			heights:  []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			name:     "example 2",
			heights:  []int{4, 3, 2, 1, 4},
			expected: 16,
		},
		{
			name:     "example 3",
			heights:  []int{1, 2, 1},
			expected: 2,
		},
		{
			name:     "example 4",
			expected: 233767415,
			heights: func() []int {
				res := make([]int, 96700)

				for i := 0; i < 96700; i++ {
					res[i] = i / 10
				}

				return res
			}(),
		},
	}
	for _, c := range cases {
		if c.name == "" {
			t.Errorf("test case name required!")
			continue
		}

		t.Run(c.name, func(t *testing.T) {
			actual := medium.MaxArea(c.heights)
			assert.Equal(t, c.expected, actual)
		})
	}
}
