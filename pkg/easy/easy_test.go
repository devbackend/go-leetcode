package easy_test

import (
	"testing"

	"github.com/devbackend/go-leetcode/pkg/easy"
	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		name     string
		num      int
		expected int
	}{
		{
			name:     "example 1",
			num:      123,
			expected: 321,
		},
		{
			name:     "example 2",
			num:      -123,
			expected: -321,
		},
		{
			name:     "example 3",
			num:      120,
			expected: 21,
		},
		{
			name:     "example 4",
			num:      0,
			expected: 0,
		},
		{
			name:     "example 5",
			num:      -4,
			expected: -4,
		},
		{
			name:     "example 6",
			num:      1534236469,
			expected: 0,
		},
	}

	for _, c := range cases {
		if c.name == "" {
			t.Errorf("test case name required!")
			continue
		}

		t.Run(c.name, func(t *testing.T) {
			actual := easy.Reverse(c.num)

			assert.Equal(t, c.expected, actual)
		})
	}
}
