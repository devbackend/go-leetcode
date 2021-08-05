package easy

import (
	"math"
	"strconv"
)

// TwoSum for https://leetcode.com/problems/two-sum/
func TwoSum(nums []int, target int) []int {
	tmp := map[int]int{}

	for k, v := range nums {
		if res, ok := tmp[v]; ok {
			return []int{res, k}
		}

		key := target - v
		tmp[key] = k
	}

	return []int{}
}

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
		tmp := chars[right]
		chars[right] = chars[left]
		chars[left] = tmp

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
