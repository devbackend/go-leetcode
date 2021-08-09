package easy_test

import (
	"testing"

	"github.com/devbackend/go-leetcode/pkg/easy"
	"github.com/stretchr/testify/assert"
)

func TestStrStr(t *testing.T) {
	cases := []struct {
		name     string
		haystack string
		needle   string
		expected int
	}{
		{
			name:     "example 1",
			haystack: "hello",
			needle:   "ll",
			expected: 2,
		},
		{
			name:     "example 2",
			haystack: "aaaaa",
			needle:   "baa",
			expected: -1,
		},
		{
			name:     "example 3",
			haystack: "",
			needle:   "",
			expected: 0,
		},
		{
			name:     "example 4",
			haystack: "abcddaaw",
			needle:   "aaaw",
			expected: -1,
		},
		{
			name:     "example 5",
			haystack: "abc",
			needle:   "abcd",
			expected: -1,
		},
	}
	for _, c := range cases {
		if c.name == "" {
			t.Errorf("test case name required!")
			continue
		}

		t.Run(c.name, func(t *testing.T) {
			actual := easy.StrStr(c.haystack, c.needle)
			assert.Equal(t, c.expected, actual)
		})
	}
}
