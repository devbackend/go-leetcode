package medium

func DetectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow := head
	fast := head

	var isCycled bool

	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next

		if fast == slow {
			isCycled = true
			break
		}
	}

	if !isCycled {
		return nil
	}

	fast = head

	for slow != fast {
		slow, fast = slow.Next, fast.Next
	}

	return fast
}
