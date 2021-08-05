package easy

// ClimbStairs for https://leetcode.com/problems/climbing-stairs
func ClimbStairs(n int) int {
	if n < 3 {
		return n
	}

	restOneStep := n - 1
	if _, ok := climbStairsCache[restOneStep]; !ok {
		climbStairsCache[restOneStep] = ClimbStairs(restOneStep)
	}

	restTwoStep := n - 2
	if _, ok := climbStairsCache[restTwoStep]; !ok {
		climbStairsCache[restTwoStep] = ClimbStairs(restTwoStep)
	}

	return climbStairsCache[restOneStep] + climbStairsCache[restTwoStep]
}
