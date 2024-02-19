package problems

var closeParentheses = map[byte]byte{
	']': '[',
	')': '(',
	'}': '{',
}

func isValid(s string) bool {
	st := make([]byte, 0, len(s))

	for ix := range s {
		par := s[ix]

		if _, ok := closeParentheses[par]; !ok {
			st = append(st, par)
			continue
		}

		popIx := len(st) - 1

		if len(st) == 0 || closeParentheses[par] != st[popIx] {
			return false
		}

		st = st[:popIx]
	}

	return len(st) == 0
}
