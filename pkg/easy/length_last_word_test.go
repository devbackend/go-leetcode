package easy_test

import (
	"testing"

	"github.com/devbackend/go-leetcode/pkg/easy"
	"github.com/stretchr/testify/assert"
)

func TestLengthOfLastWord(t *testing.T) {
	cases := []struct {
		name     string
		str      string
		expected int
	}{
		{
			name:     "example 1",
			str:      "Hello World",
			expected: 5,
		},
		{
			name:     "example 2",
			str:      "   fly me   to   the moon  ",
			expected: 4,
		},
		{
			name:     "example 3",
			str:      "luffy is still joyboy",
			expected: 6,
		},
		{
			name:     "example 4",
			str:      "luffy",
			expected: 5,
		},
		{
			name:     "example 5",
			str:      "    luffy     ",
			expected: 5,
		},
		{
			name:     "example 6",
			str:      "",
			expected: 0,
		},
		{
			name:     "example 7",
			str:      "a",
			expected: 1,
		},
	}
	for _, c := range cases {
		if c.name == "" {
			t.Errorf("test case name required!")
			continue
		}

		t.Run(c.name, func(t *testing.T) {
			actual := easy.LengthOfLastWord(c.str)
			assert.Equal(t, c.expected, actual)
		})
	}
}
