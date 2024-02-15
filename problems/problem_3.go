package problems

// 3. Longest Substring Without Repeating Characters
func lengthOfLongestSubstring(s string) int {
	chars := []byte(s)
	uniq := make(map[byte]bool, len(s))

	var l, r int

	var res int

	for r < len(chars) {
		ch := chars[r]

		for l <= r {
			if !uniq[ch] {
				break
			}

			delete(uniq, chars[l])
			l++
		}

		uniq[ch] = true

		res = max(res, len(uniq))

		r++
	}

	return res
}
