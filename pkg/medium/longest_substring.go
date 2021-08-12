package medium

import "math"

func LengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	var rights int

	chars := []byte(s)
	left := 0
	max := 1

	for left < len(s) {
		uniqueChars := map[byte]struct{}{
			chars[left]: {},
		}

		rights = left + 1
		for rights < len(s) {
			if _, ok := uniqueChars[chars[rights]]; ok {
				break
			}

			uniqueChars[chars[rights]] = struct{}{}
			rights++
			max = int(math.Max(float64(max), float64(rights-left)))
		}
		left++
	}

	return max
}
