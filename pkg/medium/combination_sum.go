package medium

// CombinationSum for https://leetcode.com/problems/combination-sum/
func CombinationSum(candidates []int, target int) [][]int {
	var res [][]int

	if target <= 0 {
		return res
	}

	for i := range candidates {
		if candidates[i] > target {
			continue
		}

		candidate := []int{candidates[i]}

		if candidates[i] == target {
			res = append(res, candidate)
			continue
		}

		prevSums := CombinationSum(candidates[i:], target-candidates[i])

		for k := range prevSums {
			res = append(res, append(candidate, prevSums[k]...))
		}
	}

	return res
}
