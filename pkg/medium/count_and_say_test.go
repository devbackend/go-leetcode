package medium_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/devbackend/go-leetcode/pkg/medium"
)

func TestCountAndSay(t *testing.T) {
	testCases := map[string]struct {
		number   int
		expected string
	}{
		"example 1": {
			number:   1,
			expected: "1",
		},
		"example 2": {
			number:   4,
			expected: "1211",
		},
		"example 3": {
			number:   2,
			expected: "11",
		},
		"example 4": {
			number:   15,
			expected: "311311222113111231131112132112311321322112111312211312111322212311322113212221",
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := medium.CountAndSay(tc.number)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
