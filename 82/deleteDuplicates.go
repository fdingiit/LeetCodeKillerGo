package deleteDuplicates

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	var dummy, pre, cur, end *ListNode

	dummy = &ListNode{Next: head}
	cur = head
	pre = dummy

	for cur != nil {
		if cur.Next != nil && cur.Next.Val == cur.Val {
			end = cur.Next
			for end.Next != nil && end.Val == end.Next.Val {
				end = end.Next
			}

			pre.Next = end.Next
			cur = end.Next
		} else {
			pre = cur
			cur = cur.Next
		}
	}

	return dummy.Next
}
