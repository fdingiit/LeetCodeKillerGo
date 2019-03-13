package partition

type ListNode struct {
	Val  int
	Next *ListNode
}

func insert(tail, target *ListNode) *ListNode {
	tail.Next = target
	tail = tail.Next
	tail.Next = nil
	return tail
}

func delete(pre, target *ListNode) *ListNode {
	if pre == nil {
		return target.Next
	}
	pre.Next = target.Next
	return pre.Next
}

func partition(head *ListNode, x int) *ListNode {
	var dummy, fixed, cur, pre, next, tail *ListNode

	dummy = &ListNode{}
	tail = dummy
	cur = head

	for cur != nil {
		if cur.Val >= x {
			if fixed == nil {
				fixed = cur
			}
			pre = cur
			cur = cur.Next
		} else {
			next = delete(pre, cur)
			tail = insert(tail, cur)
			cur = next
		}
	}

	tail.Next = fixed
	return dummy.Next
}
