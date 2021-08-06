package easy

func IsValidParentheses(s string) bool {
	chars := []byte(s)

	var opens []byte

	for _, c := range chars {
		if _, ok := parenthesesClosers[c]; !ok {
			opens = append(opens, c)
			continue
		}

		if len(opens) == 0 {
			return false
		}

		open := opens[len(opens)-1]

		if parenthesesClosers[c] != open {
			return false
		}

		opens = opens[:len(opens)-1]
	}

	return len(opens) == 0
}
