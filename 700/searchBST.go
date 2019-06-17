package searchBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	var p = root

	for {
		if p == nil {
			return nil
		}

		if p.Val == val {
			return p
		}

		if p.Val > val {
			p = p.Left
			continue
		}

		if p.Val < val {
			p = p.Right
			continue
		}
	}

	return nil
}