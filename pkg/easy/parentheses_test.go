package easy_test

import (
	"testing"

	"github.com/devbackend/go-leetcode/pkg/easy"
	"github.com/stretchr/testify/assert"
)

func TestIsValidParentheses(t *testing.T) {
	cases := []struct {
		name     string
		str      string
		expected bool
	}{
		{
			name:     "example 1",
			str:      "()",
			expected: true,
		},
		{
			name:     "example 2",
			str:      "()[]{}",
			expected: true,
		},
		{
			name:     "example 3",
			str:      "(]",
			expected: false,
		},
		{
			name:     "example 4",
			str:      "([)]",
			expected: false,
		},
		{
			name:     "example 5",
			str:      "{[]}",
			expected: true,
		},
		{
			name:     "example 6",
			str:      "}(){",
			expected: false,
		},
		{
			name:     "example 7",
			str:      "[",
			expected: false,
		},
	}
	for _, c := range cases {
		if c.name == "" {
			t.Errorf("test case name required!")
			continue
		}

		t.Run(c.name, func(t *testing.T) {
			actual := easy.IsValidParentheses(c.str)
			assert.Equal(t, c.expected, actual)
		})
	}
}
