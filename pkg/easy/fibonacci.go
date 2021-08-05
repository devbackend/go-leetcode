package easy

// Fib for https://leetcode.com/problems/fibonacci-number/
func Fib(num int) int {
	if num == 0 {
		return 0
	}

	if num < 3 {
		return 1
	}

	first := num - 1
	second := num - 2

	if _, ok := fibCache[first]; !ok {
		fibCache[first] = Fib(num - 1)
	}

	if _, ok := fibCache[second]; !ok {
		fibCache[second] = Fib(num - 2)
	}

	return fibCache[first] + fibCache[second]
}
