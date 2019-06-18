package minDepth

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	var depth int
	var curLvl, nextLvl []*TreeNode
	var p *TreeNode

	if root == nil {
		return 0
	}

	curLvl = []*TreeNode{root}

	for {
		depth++

		for len(curLvl) > 0 {
			p = curLvl[0]
			curLvl = curLvl[1:]

			if p.Left == nil && p.Right == nil {
				return depth
			}

			if p.Left != nil {
				nextLvl = append(nextLvl, p.Left)
			}

			if p.Right != nil {
				nextLvl = append(nextLvl, p.Right)
			}
		}

		if len(nextLvl) == 0 {
			return depth
		}

		curLvl = nextLvl
		nextLvl = []*TreeNode{}
	}
}
