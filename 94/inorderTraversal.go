package inorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var ret []int
	var stack = []*TreeNode{}
	var p = root

	for p != nil || len(stack) != 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}

		p = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, p.Val)

		p = p.Right
	}

	return ret
}
