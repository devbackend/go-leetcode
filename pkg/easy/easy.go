package easy

import (
	"math"
	"strconv"
)

// Reverse for https://leetcode.com/problems/reverse-integer/
func Reverse(x int) int {
	if -10 < x && x < 10 {
		return x
	}

	isNegative := x < 0

	if isNegative {
		x *= -1
	}

	chars := []byte(strconv.Itoa(x))

	left := 0
	right := len(chars) - 1

	for right > left {
		chars[right], chars[left] = chars[left], chars[right]

		left++
		right--
	}

	res, _ := strconv.Atoi(string(chars))
	if isNegative {
		res *= -1
	}

	if res < math.MinInt32 || res > math.MaxInt32 {
		return 0
	}

	return res
}
