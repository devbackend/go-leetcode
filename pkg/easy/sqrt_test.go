package easy_test

import (
	"testing"

	"github.com/devbackend/go-leetcode/pkg/easy"
	"github.com/stretchr/testify/assert"
)

func TestSqrt(t *testing.T) {
	cases := []struct {
		name     string
		num      int
		expected int
	}{
		{
			name:     "example 1",
			num:      4,
			expected: 2,
		},
		{
			name:     "example 2",
			num:      8,
			expected: 2,
		},
		{
			name:     "example 3",
			num:      16,
			expected: 4,
		},
		{
			name:     "example 4",
			num:      169,
			expected: 13,
		},
		{
			name:     "example 5",
			num:      1000000000000,
			expected: 1000000,
		},
		{
			name:     "example 6",
			num:      2,
			expected: 1,
		},
		{
			name:     "example 7",
			num:      0,
			expected: 0,
		},
		{
			name:     "example 8",
			num:      1,
			expected: 1,
		},
	}
	for _, c := range cases {
		if c.name == "" {
			t.Errorf("test case name required!")
			continue
		}

		t.Run(c.name, func(t *testing.T) {
			actual := easy.Sqrt(c.num)
			assert.Equal(t, c.expected, actual)
		})
	}
}
