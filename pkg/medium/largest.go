package medium

import (
	"sort"
	"strconv"
	"strings"
)

// LargestNumber for https://leetcode.com/problems/largest-number/
func LargestNumber(nums []int) string {
	strs := make([]string, len(nums))

	for k, v := range nums {
		strs[k] = strconv.Itoa(v)
	}

	sort.Slice(strs, func(i, j int) bool {
		first := strs[i] + strs[j]
		second := strs[j] + strs[i]

		return first > second
	})

	res := strings.Join(strs, "")
	if res[0] == '0' {
		return "0"
	}

	return res
}
