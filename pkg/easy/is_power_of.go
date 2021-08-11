package easy

// IsPowerOfTwo for https://leetcode.com/problems/power-of-two
func IsPowerOfTwo(n int) bool {
	return isPower(n, 2)
}

// IsPowerOfThree for https://leetcode.com/problems/power-of-three
func IsPowerOfThree(n int) bool {
	return isPower(n, 3)
}

// IsPowerOfFour for https://leetcode.com/problems/power-of-four
func IsPowerOfFour(n int) bool {
	return isPower(n, 4)
}

func isPower(num int, pow int) bool {
	if num <= 0 {
		return false
	}

	if num == 1 {
		return true
	}

	for num > 1 {
		if num%pow != 0 {
			return false
		}

		num /= pow
	}

	return true
}
