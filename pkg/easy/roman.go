package easy

// RomanToInt for https://leetcode.com/problems/roman-to-integer/
func RomanToInt(s string) int {
	nums := []byte(s)

	var res int

	left := 0

	for left < len(nums) {
		right := left + 1

		leftNum := romanNumbers[nums[left]]
		rightNum := 0

		if right < len(nums) {
			rightNum = romanNumbers[nums[right]]
		}

		if leftNum < rightNum {
			res += rightNum - leftNum
			left = right + 1
		} else {
			res += leftNum
			left++
		}
	}

	return res
}
