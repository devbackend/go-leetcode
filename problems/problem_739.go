package problems

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))

	decrease := make([][2]int, 0, len(temperatures))

	for i, temp := range temperatures {
		if len(decrease) == 0 {
			decrease = append(decrease, [2]int{temp, i})
			continue
		}

		for len(decrease) > 0 && decrease[len(decrease)-1][0] < temp {
			topIx := len(decrease) - 1
			tempIx := decrease[topIx][1]

			res[tempIx] = i - tempIx
			decrease = decrease[:len(decrease)-1]
		}

		decrease = append(decrease, [2]int{temp, i})
	}

	return res
}
