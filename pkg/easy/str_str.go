package easy

// StrStr for https://leetcode.com/problems/implement-strstr
func StrStr(haystack string, needle string) int {
	if haystack == needle {
		return 0
	}

	if len(needle) > len(haystack) {
		return -1
	}

	left := 0
	right := len(needle)

	for right <= len(haystack) {
		candidate := haystack[left:right]

		if candidate == needle {
			return left
		}

		left++
		right++
	}

	return -1
}
