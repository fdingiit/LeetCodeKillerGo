package levelOrder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var ret [][]int
	var cur *TreeNode
	var curLvl, nextLvl []*TreeNode

	if root == nil {
		return [][]int{}
	}

	curLvl = []*TreeNode{root}

	for {
		res := []int{}

		for len(curLvl) > 0 {
			cur = curLvl[0]
			curLvl = curLvl[1:]
			res = append(res, cur.Val)

			if cur.Left != nil {
				nextLvl = append(nextLvl, cur.Left)
			}

			if cur.Right != nil {
				nextLvl = append(nextLvl, cur.Right)
			}
		}

		ret = append(ret, res)

		if len(nextLvl) == 0 {
			return ret
		}

		curLvl = nextLvl
		nextLvl = []*TreeNode{}
	}
}
