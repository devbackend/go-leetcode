package easy

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

// TwoSumSorted for https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
func TwoSumSorted(numbers []int, target int) []int {
	nums := map[int]int{}

	for k, v := range numbers {
		if n, ok := nums[v]; ok {
			return []int{n, k + 1}
		}

		nums[target-v] = k + 1
	}

	return []int{0, 0}
}
