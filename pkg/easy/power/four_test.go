package power_test

import (
	"testing"

	"github.com/devbackend/go-leetcode/pkg/easy/power"
	"github.com/stretchr/testify/assert"
)

func TestIsPowerOfFour(t *testing.T) {
	cases := []struct {
		name     string
		num      int
		expected bool
	}{
		{
			name:     "example 1",
			num:      16,
			expected: true,
		},
		{
			name:     "example 2",
			num:      5,
			expected: false,
		},
		{
			name:     "example 3",
			num:      1,
			expected: true,
		},
		{
			name:     "example 3",
			num:      1024,
			expected: true,
		},
	}
	for _, c := range cases {
		if c.name == "" {
			t.Errorf("test case name required!")
			continue
		}

		t.Run(c.name, func(t *testing.T) {
			actual := power.IsPowerOfFour(c.num)
			assert.Equal(t, c.expected, actual)
		})
	}
}
