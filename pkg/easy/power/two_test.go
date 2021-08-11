package power_test

import (
	"testing"

	"github.com/devbackend/go-leetcode/pkg/easy/power"
	"github.com/stretchr/testify/assert"
)

func TestIsPowerOfTwo(t *testing.T) {
	cases := []struct {
		name     string
		num      int
		expected bool
	}{
		{
			name:     "example 1",
			num:      1,
			expected: true,
		},
		{
			name:     "example 2",
			num:      16,
			expected: true,
		},
		{
			name:     "example 3",
			num:      3,
			expected: false,
		},
		{
			name:     "example 4",
			num:      4,
			expected: true,
		},
		{
			name:     "example 5",
			num:      12,
			expected: false,
		},
		{
			name:     "example 6",
			num:      -8,
			expected: false,
		},
		{
			name:     "example 8",
			num:      -2047483648,
			expected: false,
		},
		{
			name:     "example 9",
			num:      0,
			expected: false,
		},
	}
	for _, c := range cases {
		if c.name == "" {
			t.Errorf("test case name required!")
			continue
		}

		t.Run(c.name, func(t *testing.T) {
			actual := power.IsPowerOfTwo(c.num)
			assert.Equal(t, c.expected, actual)
		})
	}
}
