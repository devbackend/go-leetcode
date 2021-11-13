package easy

// SearchInsert for https://leetcode.com/problems/search-insert-position/
func SearchInsert(nums []int, target int) int {
	if target <= nums[0] {
		return 0
	}

	if target > nums[len(nums)-1] {
		return len(nums)
	}

	if nums[len(nums)-1] == target {
		return len(nums) - 1
	}

	var pos int

	left := 0
	right := len(nums) - 1

	for left < right {
		if (right - left) == 1 {
			return right
		}

		k := (right + left) / 2

		if nums[k] == target {
			return k
		}

		if target < nums[k] {
			right = k
			continue
		}

		left = k
	}

	return pos
}
