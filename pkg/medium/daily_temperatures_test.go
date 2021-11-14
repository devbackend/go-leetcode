package medium_test

import (
	"testing"

	"github.com/devbackend/go-leetcode/pkg/medium"
	"github.com/stretchr/testify/assert"
)

func TestDailyTemperatures(t *testing.T) {
	cases := []struct {
		name         string
		temperatures []int
		expected     []int
	}{
		{
			name:         "example 1",
			temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73},
			expected:     []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			name:         "example 2",
			temperatures: []int{30, 40, 50, 60},
			expected:     []int{1, 1, 1, 0},
		},
		{
			name:         "example 3",
			temperatures: []int{30, 60, 90},
			expected:     []int{1, 1, 0},
		},
		{
			name:         "example 4",
			temperatures: []int{90, 60, 30},
			expected:     []int{0, 0, 0},
		},
		{
			name:         "example 5",
			temperatures: []int{60, 65, 70, 65, 60},
			expected:     []int{1, 1, 0, 0, 0},
		},
		{
			name:         "example 6",
			temperatures: []int{70, 65, 60, 65, 70},
			expected:     []int{0, 3, 1, 1, 0},
		},
		{
			name:         "example 7",
			temperatures: []int{35, 30, 32, 45},
			expected:     []int{3, 1, 1, 0},
		},
	}
	for _, c := range cases {
		if c.name == "" {
			t.Errorf("test case name required!")
			continue
		}

		t.Run(c.name, func(t *testing.T) {
			actual := medium.DailyTemperatures(c.temperatures)
			assert.Equal(t, c.expected, actual)
		})
	}
}

func BenchmarkDailyTemperatures(b *testing.B) {
	size := 9999
	temperatures := make([]int, 9999)

	for i := 0; i < size-1; i++ {
		temperatures[i] = size - i
	}

	temperatures[9998] = 10000

	for i := 0; i < b.N; i++ {
		medium.DailyTemperatures(temperatures)
	}
}
