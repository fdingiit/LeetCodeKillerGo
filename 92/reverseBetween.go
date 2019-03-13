package reverseBetween

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(head *ListNode) (*ListNode, *ListNode) {
	var pre, cur, next *ListNode

	cur = head

	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre, head
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	var cur, pre, begin, preBegin, afterEnd *ListNode
	var i int

	cur = head
	for {
		i++
		if i == m {
			preBegin = pre
			begin = cur
		}
		if i == n {
			afterEnd = cur.Next
			cur.Next = nil
			break
		}

		pre = cur
		cur = cur.Next
	}

	nHead, nTail := reverse(begin)
	if preBegin != nil {
		preBegin.Next = nHead
	}
	nTail.Next = afterEnd

	if m == 1 {
		return nHead
	} else {
		return head
	}
}