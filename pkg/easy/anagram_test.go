package easy_test

import (
	"github.com/devbackend/go-leetcode/pkg/easy"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	cases := []struct {
		name      string
		strFirst  string
		strSecond string
		expected  bool
	}{
		{
			name:      "example 1",
			strFirst:  "anagram",
			strSecond: "nagaram",
			expected:  true,
		},
		{
			name:      "example 2",
			strFirst:  "rat",
			strSecond: "car",
			expected:  false,
		},
		{
			name:      "example 3",
			strFirst:  "abc",
			strSecond: "cab",
			expected:  true,
		},
		{
			name:      "example 3",
			strFirst:  "thanos",
			strSecond: "sonata",
			expected:  false,
		},
	}

	for _, c := range cases {
		if c.name == "" {
			t.Errorf("test case name required!")
			continue
		}

		t.Run(c.name, func(t *testing.T) {
			actual := easy.IsAnagram(c.strFirst, c.strSecond)

			assert.Equal(t, c.expected, actual)
		})
	}
}
