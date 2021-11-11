package medium

// LetterCombinations for https://leetcode.com/problems/letter-combinations-of-a-phone-number/
func LetterCombinations(digits string) []string {
	letters := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	if len(digits) == 0 {
		return []string{}
	}

	syms := []byte(digits)

	if len(digits) == 1 {
		key := syms[0]
		return letters[key]
	}

	var res []string

	for _, vl := range letters[syms[0]] {
		for _, vr := range LetterCombinations(string(syms[1:])) {
			res = append(res, vl+vr)
		}
	}

	return res
}
