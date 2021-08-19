package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// HasPathSum for https://leetcode.com/problems/path-sum/
func HasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Val == targetSum && root.Right == nil && root.Left == nil {
		return true
	}

	if HasPathSum(root.Right, targetSum-root.Val) {
		return true
	}

	if HasPathSum(root.Left, targetSum-root.Val) {
		return true
	}

	return false
}
