package buildTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findIndex(arr []int, begin, end int, wanted int) int {
	for i := begin; i <= end; i++ {
		if arr[i] == wanted {
			return i
		}
	}

	return -1
}

func doBuild(preorder []int, length int, cur *int, inodrder []int, begin, end int) *TreeNode {
	if *cur >= length || begin > end {
		return nil
	}

	val := preorder[*cur]
	node := &TreeNode{Val: val}
	i := findIndex(inodrder, begin, end, val)

	*cur += 1
	node.Left = doBuild(preorder, length, cur, inodrder, begin, i-1)
	node.Right = doBuild(preorder, length, cur, inodrder, i+1, end)

	return node
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	var cur int
	return doBuild(preorder, len(preorder), &cur, inorder, 0, len(inorder)-1)
}
