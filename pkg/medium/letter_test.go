package medium_test

import (
	"testing"

	"github.com/devbackend/go-leetcode/pkg/medium"
	"github.com/stretchr/testify/assert"
)

func TestLetterCombinations(t *testing.T) {
	cases := []struct {
		name     string
		digits   string
		expected []string
	}{
		{
			name:     "example 1",
			digits:   "23",
			expected: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name:     "example 2",
			digits:   "",
			expected: []string{},
		},
		{
			name:     "example 3",
			digits:   "2",
			expected: []string{"a", "b", "c"},
		},
		{
			name:   "example 4",
			digits: "239",
			expected: []string{
				"adw", "adx", "ady", "adz",
				"aew", "aex", "aey", "aez",
				"afw", "afx", "afy", "afz",
				"bdw", "bdx", "bdy", "bdz",
				"bew", "bex", "bey", "bez",
				"bfw", "bfx", "bfy", "bfz",
				"cdw", "cdx", "cdy", "cdz",
				"cew", "cex", "cey", "cez",
				"cfw", "cfx", "cfy", "cfz",
			},
		},
	}
	for _, c := range cases {
		if c.name == "" {
			t.Errorf("test case name required!")
			continue
		}

		t.Run(c.name, func(t *testing.T) {
			actual := medium.LetterCombinations(c.digits)
			assert.ElementsMatch(t, c.expected, actual)
		})
	}
}
