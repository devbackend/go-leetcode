package problems

func productExceptSelf(nums []int) []int {
	n := len(nums)

	res := make([]int, n)
	res[n-1] = nums[n-1]

	for i := n - 2; i >= 0; i-- {
		res[i] = nums[i] * res[i+1]
	}

	left := 1
	for i := 0; i < n; i++ {
		if i == n-1 {
			res[i] = left
			continue
		}

		res[i] = left * res[i+1]
		left *= nums[i]
	}

	return res
}
