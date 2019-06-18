package kthSmallest

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorder(root *TreeNode, cnt *int, kth int) (bool, int) {
	if root.Left != nil {
		ok, val := inorder(root.Left, cnt, kth)
		if ok {
			return ok, val
		}
	}

	*cnt = *cnt + 1
	if *cnt == kth {
		return true, root.Val
	}

	if root.Right != nil {
		ok, val := inorder(root.Right, cnt, kth)
		if ok {
			return ok, val
		}
	}

	return false, -1
}

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}

	var cnt int

	_, val := inorder(root, &cnt, k)
	return val
}
