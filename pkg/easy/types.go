package easy

type ListNode struct {
	Val  int
	Next *ListNode
}

var fibCache = map[int]int{}

var romanNumbers = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

var parenthesesClosers = map[byte]byte{
	'}': '{',
	']': '[',
	')': '(',
}
