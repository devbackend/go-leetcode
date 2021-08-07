package easy

func Sqrt(x int) int {
	if x < 2 {
		return x
	}

	for i := 1; i < x; i++ {
		if i*i == x {
			return i
		}

		if i*i > x {
			return i - 1
		}
	}

	return x - 1
}
