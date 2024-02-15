package problems

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var chars [26]int

	for ix := range s {
		sCh, tCh := s[ix]-'a', t[ix]-'a'

		chars[sCh]++
		chars[tCh]--
	}

	for _, cnt := range chars {
		if cnt != 0 {
			return false
		}
	}

	return true
}
