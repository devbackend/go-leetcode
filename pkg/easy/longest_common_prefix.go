package easy

func LongestCommonPrefix(strs []string) string {
	var res string
	var ix int

	if len(strs) < 2 {
		return strs[0]
	}

	chars := make([][]byte, len(strs), len(strs))

	for k, v := range strs {
		chars[k] = []byte(v)
	}

	for ix < len(strs[0]) {
		if ix >= len(chars[0]) {
			return res
		}

		char := chars[0][ix]
		for _, v := range chars {
			if ix >= len(v) {
				return res
			}

			if v[ix] != char {
				return res
			}
		}

		res += string(char)

		ix++
	}

	return res
}
