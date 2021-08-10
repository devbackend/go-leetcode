package easy_test

import (
	"github.com/devbackend/go-leetcode/pkg/easy"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddStrings(t *testing.T) {
	cases := []struct {
		name     string
		num1     string
		num2     string
		expected string
	}{
		{
			name:     "example 1",
			num1:     "11",
			num2:     "123",
			expected: "134",
		},
		{
			name:     "example 2",
			num1:     "0",
			num2:     "0",
			expected: "0",
		},
		{
			name:     "example 3",
			num1:     "456",
			num2:     "77",
			expected: "533",
		},
		{
			name:     "example 4",
			num1:     "100000000000000000000000000011",
			num2:     "112",
			expected: "100000000000000000000000000123",
		},
		{
			name:     "example 5",
			num1:     "1",
			num2:     "9",
			expected: "10",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := easy.AddStrings(c.num1, c.num2)

			assert.Equal(t, c.expected, actual)
		})
	}
}
