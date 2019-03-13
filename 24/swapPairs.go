package swapPairs

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	var dummy, a, b, c *ListNode

	dummy = &ListNode{}
	c = dummy
	a = head

	for a != nil {
		if a.Next != nil {
			b = a.Next.Next
			c.Next = a.Next
			c = c.Next
			c.Next = a
			c = c.Next
			a = b
		} else {
			c.Next = a
			c = c.Next
			break
		}
	}

	c.Next = nil
	return dummy.Next
}
