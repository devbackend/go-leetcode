package medium

// DailyTemperatures for https://leetcode.com/problems/daily-temperatures/
func DailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))

	tempStack := make([][2]int, 0, len(temperatures))

	for k := range temperatures {
		temp := temperatures[k]

		for {
			i := len(tempStack) - 1

			if len(tempStack) == 0 || temp <= tempStack[i][1] {
				tempStack = append(tempStack, [2]int{k, temp})
				break
			}

			res[tempStack[i][0]] = k - tempStack[i][0]
			tempStack = tempStack[0:i]
		}
	}

	return res
}
