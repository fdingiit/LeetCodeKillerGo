package isValidBST

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func doCheck(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}

	if root.Val >= max || root.Val <= min {
		return false
	}

	return doCheck(root.Left, min, root.Val) && doCheck(root.Right, root.Val, max)
}

func isValidBST(root *TreeNode) bool {
	return doCheck(root, math.MinInt64, math.MaxInt64)
}
