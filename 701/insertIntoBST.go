package insertIntoBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	var (
		cur    *TreeNode
		parent *TreeNode
		insert *TreeNode
	)

	insert = &TreeNode{Val: val}
	cur = root

	for {
		if cur == nil {
			if parent == nil {
				root = insert
			} else {
				if val > parent.Val {
					parent.Right = insert
				}
				if val < parent.Val {
					parent.Left = insert
				}
			}

			break
		}

		parent = cur

		if cur.Val > val {
			cur = cur.Left
			continue
		}

		if cur.Val < val {
			cur = cur.Right
			continue
		}

	}

	return root
}
