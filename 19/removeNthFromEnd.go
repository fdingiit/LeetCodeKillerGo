package removeNthFromEnd

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var p, q *ListNode

	p = head
	q = head

	for n > 0 {
		n--
		q = q.Next
	}

	if q == nil {
		return head.Next
	}

	for q.Next != nil {
		q = q.Next
		p = p.Next
	}

	p.Next = p.Next.Next
	return head
}
