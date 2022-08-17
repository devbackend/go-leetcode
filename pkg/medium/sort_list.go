package medium

func SortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	middle := getMiddle(head)
	second := middle.Next

	middle.Next = nil

	return MergeTwoLists(SortList(head), SortList(second))
}

func getMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow := head
	fast := head.Next.Next

	for fast != nil {
		if fast.Next == nil {
			break
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func MergeTwoLists(first *ListNode, second *ListNode) *ListNode {
	if second == nil {
		return first
	}

	if first == nil {
		return second
	}

	res := &ListNode{Val: -1}
	tmp := res

	for first != nil && second != nil {
		if first.Val < second.Val {
			tmp.Next = &ListNode{Val: first.Val}
			first = first.Next
		} else {
			tmp.Next = &ListNode{Val: second.Val}
			second = second.Next
		}

		tmp = tmp.Next
	}

	if first != nil {
		tmp.Next = first
	}

	if second != nil {
		tmp.Next = second
	}

	return res.Next
}
