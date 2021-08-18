package medium

// CanJump for https://leetcode.com/problems/jump-game/
func CanJump(nums []int) bool {
	cjCache = map[int]bool{}

	return cj(0, nums)
}

var cjCache map[int]bool

func cj(start int, nums []int) bool {
	if len(nums) < 2 {
		return true
	}

	if len(nums) == 2 {
		return nums[0] > 0
	}

	if nums[0] >= len(nums)-1 {
		return true
	}

	if nums[0] == 0 {
		return false
	}

	for i := nums[0]; i > 0; i-- {
		if _, ok := cjCache[start+i]; !ok {
			cjCache[start+i] = cj(start+i, nums[i:])
		}

		if cjCache[start+i] {
			return true
		}
	}

	return false
}
