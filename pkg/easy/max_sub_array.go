package easy

import "math"

// MaxSubArray for https://leetcode.com/problems/maximum-subarray/
func MaxSubArray(nums []int) int {
	maxSum := nums[0]
	currentSum := nums[0]

	for k := 1; k < len(nums); k++ {
		currentSum = int(math.Max(float64(currentSum)+float64(nums[k]), float64(nums[k])))

		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}
