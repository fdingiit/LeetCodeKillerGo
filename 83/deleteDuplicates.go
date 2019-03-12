package deleteDuplicates

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	var ret, cur, p *ListNode

	if head == nil {
		return head
	}

	ret = head
	cur = head
	p = cur.Next

	for cur != nil  {
		if cur.Val != p.Val {
			cur.Next = p
			cur = cur.Next
		}
		p = p.Next
	}

	// NOTE:
	// be care of the tail of same val
	// eg:
	// 1, 2, 3, 3, 3
	cur.Next = nil

	return ret
}
