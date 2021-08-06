package easy_test

import (
	"testing"

	"github.com/devbackend/go-leetcode/pkg/easy"
	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	cases := []struct {
		name     string
		strings  []string
		expected string
	}{
		{
			name:     "example 1",
			strings:  []string{"flower", "flow", "flight"},
			expected: "fl",
		},
		{
			name:     "example 2",
			strings:  []string{"dog", "racecar", "car"},
			expected: "",
		},
		{
			name:     "example 3",
			strings:  []string{"", "php", "golang"},
			expected: "",
		},
		{
			name:     "example 4",
			strings:  []string{"java", "javascript", "java"},
			expected: "java",
		},
	}
	for _, c := range cases {
		if c.name == "" {
			t.Errorf("test case name required!")
			continue
		}

		t.Run(c.name, func(t *testing.T) {
			actual := easy.LongestCommonPrefix(c.strings)
			assert.Equal(t, c.expected, actual)
		})
	}
}
