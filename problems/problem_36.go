package problems

func isValidSudoku(board [][]byte) bool {
	var horizontal, vertical, trinity [9][9]bool

	for r := range board {
		for c := range board[r] {
			val := board[r][c]
			if val == '.' {
				continue
			}

			val -= '1'

			trinityPos := (r / 3) + (c/3)*3

			if horizontal[r][val] ||
				vertical[c][val] ||
				trinity[trinityPos][val] {
				return false
			}

			horizontal[r][val] = true
			vertical[c][val] = true
			trinity[trinityPos][val] = true
		}
	}

	return true
}
