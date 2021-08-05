package easy_test

import (
	"testing"

	"github.com/devbackend/go-leetcode/pkg/easy"
	"github.com/stretchr/testify/assert"
)

func TestFib(t *testing.T) {
	cases := []struct {
		name     string
		num      int
		expected int
	}{
		{
			name:     "fib by zero",
			num:      0,
			expected: 0,
		},
		{
			name:     "First number",
			num:      1,
			expected: 1,
		},
		{
			name:     "Second number",
			num:      2,
			expected: 1,
		},
		{
			name:     "Third number",
			num:      3,
			expected: 2,
		},
		{
			name:     "Third number",
			num:      4,
			expected: 3,
		},
		{
			name:     "Third number",
			num:      50,
			expected: 12586269025,
		},
	}
	for _, c := range cases {
		if c.name == "" {
			t.Errorf("test case name required!")
			continue
		}

		t.Run(c.name, func(t *testing.T) {
			actual := easy.Fib(c.num)

			assert.Equal(t, c.expected, actual)
		})
	}
}
