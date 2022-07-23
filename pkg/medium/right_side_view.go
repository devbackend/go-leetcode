package medium

func RightSideView(root *TreeNode) []int {
	var res []int

	var max int

	rightSideViewRec(root, 1, &res, &max)

	return res
}

func rightSideViewRec(root *TreeNode, level int, res *[]int, maxLevel *int) {
	if root == nil {
		return
	}

	if *maxLevel < level {
		*res = append(*res, root.Val)
		*maxLevel = level
	}

	rightSideViewRec(root.Right, level+1, res, maxLevel)
	rightSideViewRec(root.Left, level+1, res, maxLevel)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
