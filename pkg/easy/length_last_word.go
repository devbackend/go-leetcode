package easy

func LengthOfLastWord(s string) int {
	var res int

	chars := []byte(s)

	if len(chars) == 0 {
		return 0
	}

	k := len(chars)

	var startCount bool

	for k > 0 {
		k--

		if chars[k] == ' ' {
			if !startCount {
				continue
			}

			break
		}

		if !startCount {
			startCount = true
		}

		res++
	}

	return res
}
