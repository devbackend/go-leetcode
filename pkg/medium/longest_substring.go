package medium

func LengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	var right int

	chars := []byte(s)
	left := 0
	max := 1

	uniqueChars := map[byte]int{}

	for right < len(s) {
		if pos, ok := uniqueChars[chars[right]]; ok && pos >= left {
			left = pos + 1
		}

		uniqueChars[chars[left]] = left
		uniqueChars[chars[right]] = right
		right++

		diff := right - left
		if diff > max {
			max = diff
		}
	}

	return max
}
