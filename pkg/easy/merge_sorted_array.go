package easy

// Merge for https://leetcode.com/problems/merge-sorted-array
func Merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	if m == 0 {
		for k, v := range nums2 {
			nums1[k] = v
		}

		return
	}

	var first, second int

	for second < len(nums2) {
		if first >= m {
			nums1[first] = nums2[second]
			first++
			second++
			continue
		}

		if nums1[first] > nums2[second] {
			for k := m; k >= first+1; k-- {
				nums1[k] = nums1[k-1]
			}

			nums1[first] = nums2[second]
			second++
			m++
		}

		first++
	}
}
