package problems

func maxArea(heights []int) int {
	l, r := 0, len(heights)-1

	var res int
	for l < r {
		h1, h2 := heights[l], heights[r]

		res = max(res, (r-l)*min(h1, h2))

		if h1 < h2 {
			l++
		} else {
			r--
		}
	}

	return res
}
