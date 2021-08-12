package medium

import "math"

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
		max = int(math.Max(float64(max), float64(right-left)))
	}

	return max
}
