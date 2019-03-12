package isPalindrome

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre, cur, next *ListNode

	cur = head

	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}

func isPalindrome(head *ListNode) bool {
	var p, q *ListNode
	var cnt, i int

	p = head
	for p != nil {
		p = p.Next
		cnt++
	}

	i = cnt/2 + cnt%2
	q = head

	for i > 0 {
		q = q.Next
		i--
	}

	q = reverseList(q)

	p = head
	for q != nil {
		if p.Val != q.Val {
			return false
		}
		p = p.Next
		q = q.Next
	}

	return true
}
