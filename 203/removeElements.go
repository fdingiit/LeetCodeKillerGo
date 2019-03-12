package removeElements

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	var dummy, pre, cur *ListNode

	dummy = &ListNode{Next: head}
	cur = head
	pre = dummy

	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		cur = cur.Next
	}

	return dummy.Next
}
