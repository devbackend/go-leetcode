package medium_test

import (
	"testing"

	"github.com/devbackend/go-leetcode/pkg/medium"
	"github.com/stretchr/testify/assert"
)

func TestCanJump(t *testing.T) {
	cases := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "example 1",
			nums:     []int{2, 3, 1, 1, 4},
			expected: true,
		},
		{
			name:     "example 2",
			nums:     []int{3, 2, 1, 0, 4},
			expected: false,
		},
		{
			name:     "example 3",
			nums:     []int{0},
			expected: true,
		},
		{
			name:     "example 4",
			nums:     []int{1, 0},
			expected: true,
		},
		{
			name:     "example 5",
			nums:     []int{0, 1},
			expected: false,
		},
		{
			name:     "example 6",
			nums:     []int{3, 2, 2, 0, 4},
			expected: true,
		},
		{
			name:     "example 7",
			nums:     []int{5, 2, 2, 3, 0, 1},
			expected: true,
		},
		{
			name:     "example 8",
			nums:     []int{0, 2, 2, 3, 0, 1},
			expected: false,
		},
		{
			name:     "example 9",
			nums:     []int{1, 2, 3},
			expected: true,
		},
	}
	for _, c := range cases {
		if c.name == "" {
			t.Errorf("test case name required!")
			continue
		}

		t.Run(c.name, func(t *testing.T) {
			actual := medium.CanJump(c.nums)
			assert.Equal(t, c.expected, actual)
		})
	}
}

func BenchmarkCanJump(b *testing.B) {
	benchmarks := []struct {
		name string
		nums []int
	}{
		{
			name: "example 1",
			nums: []int{2, 3, 1, 1, 4},
		},
		{
			name: "example 2",
			nums: []int{3, 2, 1, 0, 4},
		},
		{
			name: "example 3",
			nums: func() []int {
				res := make([]int, 9500)
				for i := 0; i < 9500; i++ {
					if i >= 9000 {
						res[i] = 1
					} else {
						res[i] = 9000 - i
					}
				}
				return res
			}(),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				medium.CanJump(bm.nums)
			}
		})
	}
}
