package problems

func containsDuplicate(nums []int) bool {
	exists := make(map[int]bool)

	for _, num := range nums {
		if exists[num] {
			return true
		}
		exists[num] = true
	}
	return false
}
