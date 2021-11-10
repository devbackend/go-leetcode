package medium

// GenerateParenthesis for https://leetcode.com/problems/generate-parentheses/
func GenerateParenthesis(n int) []string {
	var res []string

	if n == 0 {
		return []string{""}
	}

	for i := 0; i < n; i++ {
		for _, left := range GenerateParenthesis(i) {
			for _, right := range GenerateParenthesis(n - i - 1) {
				res = append(res, "("+left+")"+right)
			}
		}
	}

	return res
}
