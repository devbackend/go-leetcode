package problems

func twoSum(nums []int, target int) []int {
	buf := make(map[int]int, len(nums))

	for i, num := range nums {
		if ix, ok := buf[target-num]; ok {
			return []int{ix, i}
		}

		buf[num] = i
	}

	return nil
}
