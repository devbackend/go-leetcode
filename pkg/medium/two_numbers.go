package medium

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{
		Val: (l1.Val + l2.Val) % 10,
	}
	curr := head

	sum := (l1.Val + l2.Val) / 10

	first := l1.Next
	second := l2.Next

	for first != nil || second != nil {
		if first != nil {
			sum += first.Val
			first = first.Next
		}

		if second != nil {
			sum += second.Val
			second = second.Next
		}

		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next

		sum /= 10
	}

	if sum > 0 {
		curr.Next = &ListNode{Val: sum % 10}
	}

	return head
}
