package easy_test

import (
	"testing"

	"github.com/devbackend/go-leetcode/pkg/easy"
	"github.com/stretchr/testify/assert"
)

func TestIsPowerOfTwo(t *testing.T) {
	cases := []struct {
		name     string
		num      int
		expected bool
	}{
		{
			name:     "example 1",
			num:      1,
			expected: true,
		},
		{
			name:     "example 2",
			num:      16,
			expected: true,
		},
		{
			name:     "example 3",
			num:      3,
			expected: false,
		},
		{
			name:     "example 4",
			num:      4,
			expected: true,
		},
		{
			name:     "example 5",
			num:      12,
			expected: false,
		},
		{
			name:     "example 6",
			num:      -8,
			expected: false,
		},
		{
			name:     "example 8",
			num:      -2047483648,
			expected: false,
		},
		{
			name:     "example 9",
			num:      0,
			expected: false,
		},
	}
	for _, c := range cases {
		if c.name == "" {
			t.Errorf("test case name required!")
			continue
		}

		t.Run(c.name, func(t *testing.T) {
			actual := easy.IsPowerOfTwo(c.num)
			assert.Equal(t, c.expected, actual)
		})
	}
}

func TestIsPowerOfThree(t *testing.T) {
	cases := []struct {
		name     string
		num      int
		expected bool
	}{
		{
			name:     "example 1",
			num:      27,
			expected: true,
		},
		{
			name:     "example 2",
			num:      0,
			expected: false,
		},
		{
			name:     "example 3",
			num:      9,
			expected: true,
		},
		{
			name:     "example 4",
			num:      45,
			expected: false,
		},
		{
			name:     "example 5",
			num:      -9,
			expected: false,
		},
		{
			name:     "example 6",
			num:      1,
			expected: true,
		},
	}
	for _, c := range cases {
		if c.name == "" {
			t.Errorf("test case name required!")
			continue
		}

		t.Run(c.name, func(t *testing.T) {
			actual := easy.IsPowerOfThree(c.num)
			assert.Equal(t, c.expected, actual)
		})
	}
}

func TestIsPowerOfFour(t *testing.T) {
	cases := []struct {
		name     string
		num      int
		expected bool
	}{
		{
			name:     "example 1",
			num:      16,
			expected: true,
		},
		{
			name:     "example 2",
			num:      5,
			expected: false,
		},
		{
			name:     "example 3",
			num:      1,
			expected: true,
		},
		{
			name:     "example 3",
			num:      1024,
			expected: true,
		},
	}
	for _, c := range cases {
		if c.name == "" {
			t.Errorf("test case name required!")
			continue
		}

		t.Run(c.name, func(t *testing.T) {
			actual := easy.IsPowerOfFour(c.num)
			assert.Equal(t, c.expected, actual)
		})
	}
}
