package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProblem_isValid(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		str      string
		expected bool
	}{
		"example 1": {
			str:      "()",
			expected: true,
		},
		"example 2": {
			str:      "()[]{}",
			expected: true,
		},
		"example 3": {
			str:      "(]",
			expected: false,
		},
		"example 4": {
			str:      "([)]",
			expected: false,
		},
		"example 5": {
			str:      "{[]}",
			expected: true,
		},
		"example 6": {
			str:      "}(){",
			expected: false,
		},
		"example 7": {
			str:      "[",
			expected: false,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			require.Equal(t, tc.expected, isValid(tc.str))
		})
	}
}
