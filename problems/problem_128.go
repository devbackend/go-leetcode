package problems

func longestConsecutive(nums []int) int {
	unique := make(map[int]bool, len(nums))
	for _, num := range nums {
		unique[num] = true
	}

	var res int

	for num := range unique {
		prev := num - 1
		if unique[prev] {
			continue
		}

		chain := num + 1
		for unique[chain] {
			chain++
		}

		res = max(res, chain-num)
	}

	return res
}
