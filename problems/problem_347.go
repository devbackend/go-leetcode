package problems

func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]int, len(nums))
	for _, num := range nums {
		counts[num]++
	}

	cells := make([][]int, len(nums)+1)
	for num, count := range counts {
		cells[count] = append(cells[count], num)
	}

	res := make([]int, 0, k)

	ix := len(cells)

	for len(res) < k {
		ix--

		for _, num := range cells[ix] {
			res = append(res, num)
			if len(res) == k {
				break
			}
		}
	}

	return res
}
