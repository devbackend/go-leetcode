package medium

// IsValidSudoku for https://leetcode.com/problems/valid-sudoku/
func IsValidSudoku(board [][]byte) bool {
	horizontal := make([]map[byte]struct{}, 9)
	vertical := make([]map[byte]struct{}, 9)
	trinity := make([]map[byte]struct{}, 9)

	for k := range board {
		for i := range board[k] {
			val := board[k][i]

			if val == '.' {
				continue
			}

			trinityPos := (i/3)*3 + (k / 3)

			if horizontal[k] == nil {
				horizontal[k] = map[byte]struct{}{}
			}

			if vertical[i] == nil {
				vertical[i] = map[byte]struct{}{}
			}

			if trinity[trinityPos] == nil {
				trinity[trinityPos] = map[byte]struct{}{}
			}

			if _, ok := horizontal[k][val]; ok {
				return false
			}

			horizontal[k][val] = struct{}{}

			if _, ok := vertical[i][val]; ok {
				return false
			}

			vertical[i][val] = struct{}{}

			if _, ok := trinity[trinityPos][val]; ok {
				return false
			}

			trinity[trinityPos][val] = struct{}{}
		}
	}

	return true
}
