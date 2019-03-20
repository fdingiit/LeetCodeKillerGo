package isUnivalTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var p *TreeNode
	n := root.Val
	stack := []*TreeNode{root}

	for len(stack) != 0 {
		p = stack[0]
		if p.Left != nil {
			stack = append(stack, p.Left)
		}
		if p.Right != nil {
			stack = append(stack, p.Right)
		}
		stack = stack[1:]

		if n != p.Val {
			return false
		}
	}

	return true
}
