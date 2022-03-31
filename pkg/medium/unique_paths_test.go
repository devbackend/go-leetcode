package medium_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/devbackend/go-leetcode/pkg/medium"
)

func TestUniquePaths(t *testing.T) {
	testCases := map[string]struct {
		rows     int
		columns  int
		expected int
	}{
		"example 1": {
			rows:     3,
			columns:  7,
			expected: 28,
		},
		"example 2": {
			rows:     3,
			columns:  2,
			expected: 3,
		},
		"example 3": {
			rows:     51,
			columns:  9,
			expected: 1916797311,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := medium.UniquePaths(tc.rows, tc.columns)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
