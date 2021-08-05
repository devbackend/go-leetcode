package easy

type ListNode struct {
	Val  int
	Next *ListNode
}

var climbStairsCache = make(map[int]int)
