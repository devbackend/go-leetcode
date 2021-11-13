package medium

import (
	"sort"
	"strconv"
)

// ThreeSum for https://leetcode.com/problems/3sum/
func ThreeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	buf := make(map[int]int)

	for k := range nums {
		buf[0-nums[k]] = k
	}

	var res [][]int

	exists := make(map[string]struct{})

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]

			k, ok := buf[sum]
			if !ok || k == i || k == j {
				continue
			}

			variant := []int{nums[i], nums[j], nums[k]}

			sort.Ints(variant)

			key := strconv.Itoa(variant[0]) + "|" + strconv.Itoa(variant[1]) + "|" + strconv.Itoa(variant[2])
			if _, ok = exists[key]; ok {
				continue
			}

			exists[key] = struct{}{}

			res = append(res, variant)
		}
	}

	return res
}
