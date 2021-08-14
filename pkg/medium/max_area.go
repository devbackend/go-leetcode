package medium

func MaxArea(height []int) int {
	max := 0

	left := 0
	right := len(height) - 1

	var leftVal, rightVal int

	for left < right {
		leftVal = height[left]
		rightVal = height[right]

		minHeight := leftVal
		if rightVal < minHeight {
			minHeight = rightVal
		}

		square := (right - left) * minHeight
		if square > max {
			max = square
		}

		if leftVal > rightVal {
			right--
		} else {
			left++
		}
	}

	return max
}
