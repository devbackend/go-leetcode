package problems

func groupAnagrams(strs []string) [][]string {
	type similarity [26]int

	groups := make(map[similarity][]string, len(strs))
	for _, word := range strs {
		var key similarity

		for _, ch := range word {
			key[ch-'a']++
		}

		groups[key] = append(groups[key], word)
	}

	res := make([][]string, 0, len(groups))

	for _, words := range groups {
		res = append(res, words)
	}

	return res
}
