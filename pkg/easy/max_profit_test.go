package easy_test

import (
	"testing"

	"github.com/devbackend/go-leetcode/pkg/easy"
	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	testCases := map[string]struct {
		prices   []int
		expected int
	}{
		"example 1": {
			prices:   []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		"example 2": {
			prices:   []int{7, 6, 4, 3, 1},
			expected: 0,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := easy.MaxProfit(tc.prices)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
