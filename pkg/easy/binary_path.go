package easy

import (
	"fmt"
	"strconv"
)

// BinaryTreePaths for https://leetcode.com/problems/binary-tree-paths/
func BinaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	left := BinaryTreePaths(root.Left)
	right := BinaryTreePaths(root.Right)

	head := strconv.Itoa(root.Val)

	if len(left) == 0 && len(right) == 0 {
		return []string{head}
	}

	res := make([]string, 0, len(left)+len(right))

	for k := range left {
		res = append(res, fmt.Sprintf("%s->%s", head, left[k]))
	}

	for k := range right {
		res = append(res, fmt.Sprintf("%s->%s", head, right[k]))
	}

	return res
}
