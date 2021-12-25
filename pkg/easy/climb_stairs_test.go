package easy_test

import (
	"testing"

	"github.com/devbackend/go-leetcode/pkg/easy"
	"github.com/stretchr/testify/assert"
)

func TestClimbStairs(t *testing.T) {
	cases := []struct {
		name     string
		stairs   int
		expected int
	}{
		{
			name:     "Example 1",
			stairs:   2,
			expected: 2,
		},
		{
			name:     "Example 2",
			stairs:   3,
			expected: 3,
		},
		{
			name:     "Example 3",
			stairs:   4,
			expected: 5,
		},
		{
			name:     "Example 4",
			stairs:   5,
			expected: 8,
		},
		{
			name:     "Example 5",
			stairs:   42,
			expected: 433494437,
		},
	}
	for _, c := range cases {
		if c.name == "" {
			t.Errorf("test case name required!")
			continue
		}

		t.Run(c.name, func(t *testing.T) {
			actual := easy.ClimbStairs(c.stairs)

			assert.Equal(t, c.expected, actual)
		})
	}
}
