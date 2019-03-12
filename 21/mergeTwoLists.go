package mergeTwoLists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var dummy, cur, p, q *ListNode

	dummy = &ListNode{}
	cur = dummy
	p = l1
	q = l2

	for p != nil && q != nil {
		if p.Val <= q.Val {
			cur.Next = p
			p = p.Next
		} else {
			cur.Next = q
			q = q.Next
		}
		cur = cur.Next
	}

	if p != nil {
		cur.Next = p
	}
	if q != nil {
		cur.Next = q
	}

	return dummy.Next
}
