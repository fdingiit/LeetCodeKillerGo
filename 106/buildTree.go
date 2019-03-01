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

func doBuild(postorder []int, cur *int, inorder []int, begin, end int) *TreeNode {
	if *cur < 0 || begin > end {
		return nil
	}

	val := postorder[*cur]
	node := &TreeNode{Val: val}
	i := findIndex(inorder, begin, end, val)

	*cur -= 1
	node.Right = doBuild(postorder, cur, inorder, i+1, end)
	node.Left = doBuild(postorder, cur, inorder, begin, i-1)

	return node
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	var cur = len(postorder) - 1
	return doBuild(postorder, &cur, inorder, 0, len(inorder)-1)
}
