package invertTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var p *TreeNode

	p = root.Left
	root.Left = root.Right
	root.Right = p

	invertTree(root.Left)
	invertTree(root.Right)

	return root
}
