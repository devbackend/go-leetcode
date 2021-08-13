package medium

func MaxArea(height []int) int {
	max := 0

	for i := 0; i < len(height); i++ {
		for k := i + 1; k < len(height); k++ {
			minHeight := height[i]
			if height[k] < minHeight {
				minHeight = height[k]
			}

			square := (k - i) * minHeight
			if square > max {
				max = square
			}
		}
	}

	return max
}
