package listnode

// Reverse for https://leetcode.com/problems/reverse-linked-list/
func Reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	res := Reverse(head.Next)

	head.Next.Next = head
	head.Next = nil

	return res
}
