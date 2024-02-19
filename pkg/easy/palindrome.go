package easy

import (
	"strconv"
)

// IsPalindromeNumber for https://leetcode.com/problems/palindrome-linked-list/
func IsPalindromeNumber(x int) bool {
	if 0 <= x && x <= 9 {
		return true
	}

	chars := []byte(strconv.Itoa(x))

	left := 0
	right := len(chars) - 1

	for right > left {
		if chars[left] != chars[right] {
			return false
		}

		left++
		right--
	}

	return true
}

// IsPalindromeLinkedList for https://leetcode.com/problems/palindrome-linked-list/
func IsPalindromeLinkedList(head *ListNode) bool {
	node := head

	var nums []int
	for node != nil {
		nums = append(nums, node.Val)

		node = node.Next
	}

	left := 0
	right := len(nums) - 1

	for right > left {
		if nums[left] != nums[right] {
			return false
		}

		left++
		right--
	}

	return true
}
