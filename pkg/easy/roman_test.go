package easy_test

import (
	"testing"

	"github.com/devbackend/go-leetcode/pkg/easy"
	"github.com/stretchr/testify/assert"
)

func TestRomanToInt(t *testing.T) {
	cases := []struct {
		name     string
		roman    string
		expected int
	}{
		{
			name:     "example 1",
			roman:    "III",
			expected: 3,
		},
		{
			name:     "example 2",
			roman:    "IV",
			expected: 4,
		},
		{
			name:     "example 3",
			roman:    "IX",
			expected: 9,
		},
		{
			name:     "example 4",
			roman:    "LVIII",
			expected: 58,
		},
		{
			name:     "example 5",
			roman:    "MCMXCIV",
			expected: 1994,
		},
	}
	for _, c := range cases {
		if c.name == "" {
			t.Errorf("test case name required!")
			continue
		}

		t.Run(c.name, func(t *testing.T) {
			actual := easy.RomanToInt(c.roman)
			assert.Equal(t, c.expected, actual)
		})
	}
}
