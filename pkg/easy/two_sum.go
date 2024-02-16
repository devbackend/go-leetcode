package easy

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
