package listnode

// BuildFromSlice for create linked list from slice
func BuildFromSlice(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}

	curr := head
	for i := 1; i < len(nums); i++ {
		curr.Next = &ListNode{Val: nums[i]}

		curr = curr.Next
	}

	return head
}
