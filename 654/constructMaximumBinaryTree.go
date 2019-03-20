package constructMaximumBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMaxIndex(nums []int, begin, end int) (index int) {
	if begin > end {
		return -1
	}

	index = begin
	max := nums[begin]

	for i := begin; i <= end; i++ {
		if nums[i] > max {
			index = i
			max = nums[i]
		}
	}

	return index
}

func doConstruct(nums []int, begin, end int) *TreeNode {
	i := findMaxIndex(nums, begin, end)

	if i == -1 {
		return nil
	}

	root := &TreeNode{Val: nums[i]}
	root.Left = doConstruct(nums, begin, i-1)
	root.Right = doConstruct(nums, i+1, end)

	return root
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return doConstruct(nums, 0, len(nums)-1)
}