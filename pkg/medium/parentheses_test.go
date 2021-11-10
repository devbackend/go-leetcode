package medium_test

import (
	"sort"
	"testing"

	"github.com/devbackend/go-leetcode/pkg/medium"
	"github.com/stretchr/testify/assert"
)

func TestGenerateParenthesis(t *testing.T) {
	cases := []struct {
		name     string
		num      int
		expected []string
	}{
		{
			name:     "example 1",
			num:      3,
			expected: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			name:     "example 2",
			num:      1,
			expected: []string{"()"},
		},
		{
			name:     "example 3",
			num:      2,
			expected: []string{"()()", "(())"},
		},
		{
			name: "example 4",
			num:  4,
			expected: []string{
				"()(())()",
				"()(()())",
				"()()(())",
				"(()(()))",
				"(())()()",
				"((())())",
				"()((()))",
				"((()))()",
				"(((())))",
				"()()()()",
				"(()()())",
				"(()())()",
				"((()()))",
				"(())(())",
			},
		},
	}
	for _, c := range cases {
		if c.name == "" {
			t.Errorf("test case name required!")
			continue
		}

		t.Run(c.name, func(t *testing.T) {
			actual := medium.GenerateParenthesis(c.num)

			sort.Strings(actual)
			sort.Strings(c.expected)

			assert.Equal(t, c.expected, actual)
		})
	}
}
