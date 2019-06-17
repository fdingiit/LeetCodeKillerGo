package largestValues

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	var (
		curLevel  []*TreeNode
		nextLevel []*TreeNode
		p         *TreeNode
		max       int
		ret       []int
	)

	if root == nil {
		return []int{}
	}

	curLevel = []*TreeNode{root}

	for {
		if len(curLevel) == 0 {
			break
		}

		max = curLevel[0].Val
		nextLevel = []*TreeNode{}

		for len(curLevel) != 0 {
			p = curLevel[0]
			curLevel = curLevel[1:]

			if p.Val > max {
				max = p.Val
			}

			if p.Left != nil {
				nextLevel = append(nextLevel, p.Left)
			}

			if p.Right != nil {
				nextLevel = append(nextLevel, p.Right)
			}
		}

		ret = append(ret, max)
		curLevel = nextLevel
	}

	return ret
}
