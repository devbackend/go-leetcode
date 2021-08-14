package easy

// RemoveDuplicates for https://leetcode.com/problems/remove-duplicates-from-sorted-array/
func RemoveDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	if nums[0] == nums[len(nums)-1] {
		return 1
	}

	diff := nums[len(nums)-1] - nums[0]

	if diff == 1 {
		nums[1] = nums[len(nums)-1]
		return 2
	}

	prev := nums[0]
	pos := 1
	uniqueCount := 1

	for _, v := range nums {
		if v == prev {
			continue
		}

		nums[pos] = v
		prev = v

		pos++
		uniqueCount++
	}

	return uniqueCount
}
