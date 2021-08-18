package easy

// SingleNumber for https://leetcode.com/problems/single-number/
func SingleNumber(nums []int) int {
	doubles := map[int]struct{}{}

	res := 0

	for _, v := range nums {
		if _, ok := doubles[v]; !ok {
			doubles[v] = struct{}{}
			res += v

			continue
		}

		res -= v
	}

	return res
}
