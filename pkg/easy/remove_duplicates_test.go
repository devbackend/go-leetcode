package easy_test

import (
	"testing"

	"github.com/devbackend/go-leetcode/pkg/easy"
	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	cases := []struct {
		name         string
		nums         []int
		expectedNums []int
	}{
		{
			name:         "example 1",
			nums:         []int{1, 1, 2},
			expectedNums: []int{1, 2},
		},
		{
			name:         "example 2",
			nums:         []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expectedNums: []int{0, 1, 2, 3, 4},
		},
		{
			name:         "example 3",
			nums:         []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4, 4, 4, 4, 4},
			expectedNums: []int{0, 1, 2, 3, 4},
		},
		{
			name:         "example 5",
			nums:         []int{},
			expectedNums: []int{},
		},
		{
			name:         "example 6",
			nums:         []int{0, 0, 0},
			expectedNums: []int{0},
		},
		{
			name:         "example 7",
			nums:         []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			expectedNums: []int{0},
		},
		{
			name:         "example 8",
			nums:         []int{0},
			expectedNums: []int{0},
		},
		{
			name:         "example 9",
			nums:         []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			expectedNums: []int{0, 1},
		},
	}
	for _, c := range cases {
		if c.name == "" {
			t.Errorf("test case name required!")
			continue
		}

		t.Run(c.name, func(t *testing.T) {
			actual := easy.RemoveDuplicates(c.nums)
			assert.Equal(t, len(c.expectedNums), actual)

			for i := 0; i < len(c.expectedNums); i++ {
				assert.Equal(t, c.expectedNums[i], c.nums[i])
			}
		})
	}
}

func BenchmarkRemoveDuplicates(b *testing.B) {
	benchmarks := []struct {
		name string
		nums []int
	}{
		{
			name: "unique only",
			nums: func() []int {
				nums := make([]int, 30000)

				for i := 0; i < len(nums); i++ {
					nums[i] = i
				}

				return nums
			}(),
		},
		{
			name: "all same",
			nums: func() []int {
				nums := make([]int, 30000)

				for i := 0; i < len(nums); i++ {
					nums[i] = 1
				}

				return nums
			}(),
		},
		{
			name: "example from leetcode",
			nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				easy.RemoveDuplicates(bm.nums)
			}
		})
	}
}
