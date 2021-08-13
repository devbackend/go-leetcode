package easy

// PlusOne for https://leetcode.com/problems/plus-one
func PlusOne(digits []int) []int {
	res := make([]int, len(digits)+1)

	diff := 1

	pos := len(digits) - 1

	for pos >= 0 {
		res[pos] = (digits[pos] + diff) % 10

		diff = (digits[pos] + diff) / 10

		pos--
	}

	if diff == 0 {
		res = res[0 : len(res)-1]
	} else {
		for k := 0; k < len(res)-1; k++ {
			res[k+1] = res[k]
		}

		res[0] = diff
	}

	return res
}
