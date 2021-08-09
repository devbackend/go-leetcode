package medium

// LongestPalindrome for https://leetcode.com/problems/longest-palindromic-substring/
func LongestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	ws := 0

	for ws < len(s) {
		start := 0
		finish := len(s) - ws

		for finish <= len(s) {
			chars := []byte(s)[start:finish]

			l := 0
			r := len(chars) - 1

			isPalindrome := true

			for r > l {
				if chars[l] != chars[r] {
					isPalindrome = false
					break
				}

				l++
				r--
			}

			if isPalindrome {
				return string(chars)
			}

			start++
			finish++
		}

		ws++
	}

	return s[0:0]
}
