package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProblemDailyTemperatures(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		temperatures []int
		expected     []int
	}{
		"example 1": {
			temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73},
			expected:     []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		"example 2": {
			temperatures: []int{30, 40, 50, 60},
			expected:     []int{1, 1, 1, 0},
		},
		"example 3": {
			temperatures: []int{30, 60, 90},
			expected:     []int{1, 1, 0},
		},
		"example 4": {
			temperatures: []int{90, 60, 30},
			expected:     []int{0, 0, 0},
		},
		"example 5": {
			temperatures: []int{60, 65, 70, 65, 60},
			expected:     []int{1, 1, 0, 0, 0},
		},
		"example 6": {
			temperatures: []int{70, 65, 60, 65, 70},
			expected:     []int{0, 3, 1, 1, 0},
		},
		"example 7": {
			temperatures: []int{35, 30, 32, 45},
			expected:     []int{3, 1, 1, 0},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			require.Equal(t, tc.expected, dailyTemperatures(tc.temperatures))
		})
	}
}
