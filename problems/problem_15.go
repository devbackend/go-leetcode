package problems

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	var res [][]int

	for i, num := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1

		for l < r {
			left, right := nums[l], nums[r]

			sum := num + left + right

			if sum == 0 {
				res = append(res, []int{num, left, right})

				l++

				for nums[l] == nums[l-1] && l < r {
					l++
				}

				continue
			}

			if sum > 0 {
				r--
			} else {
				l++
			}
		}
	}

	return res
}
