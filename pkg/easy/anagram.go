package easy

// IsAnagram for https://leetcode.com/problems/valid-anagram/
func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	buffer := map[byte]int{}

	first := []byte(s)
	second := []byte(t)

	for k, v := range first {
		if _, ok := buffer[v]; !ok {
			buffer[v] = 0
		}

		buffer[v]++

		if _, ok := buffer[second[k]]; !ok {
			buffer[second[k]] = 0
		}

		buffer[second[k]]--
	}

	for _, v := range buffer {
		if v != 0 {
			return false
		}
	}

	return true
}
