package easy

// HasCycle for https://leetcode.com/problems/linked-list-cycle/
func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next.Next

	for slow != nil && fast != nil {
		if fast == slow {
			return true
		}

		if fast.Next == nil {
			return false
		}

		slow, fast = slow.Next, fast.Next.Next
	}

	return false
}
