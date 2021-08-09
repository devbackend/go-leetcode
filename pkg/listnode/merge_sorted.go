package listnode

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	first := l1
	second := l2

	var head *ListNode
	if first.Val < second.Val {
		head = first
		first = first.Next
	} else {
		head = second
		second = second.Next
	}

	curr := head

	for first != nil && second != nil {
		if first.Val < second.Val {
			curr.Next = first
			curr = curr.Next
			first = first.Next
		} else {
			curr.Next = second
			curr = curr.Next
			second = second.Next
		}
	}

	if first != nil {
		curr.Next = first
	}

	if second != nil {
		curr.Next = second
	}

	return head
}
