package medium

// UniquePaths for https://leetcode.com/problems/unique-paths/
func UniquePaths(m int, n int) int {
	mem := make([][]int, m)

	for i := range mem {
		mem[i] = make([]int, n)
	}

	return memHelper(m, n, mem)
}

func memHelper(m int, n int, mem [][]int) int {
	if m == 2 && n == 2 {
		return 2
	}

	if m == 1 || n == 1 {
		return 1
	}

	if mem[m-2][n-1] == 0 {
		mem[m-2][n-1] = memHelper(m-1, n, mem)
	}

	if mem[m-1][n-2] == 0 {
		mem[m-1][n-2] = memHelper(m, n-1, mem)
	}

	return mem[m-2][n-1] + mem[m-1][n-2]
}
