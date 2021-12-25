package easy

// SingleNumber for https://leetcode.com/problems/single-number/
func SingleNumber(nums []int) int {
	res := 0

	for _, v := range nums {
		res ^= v
	}

	return res
}
