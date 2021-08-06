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