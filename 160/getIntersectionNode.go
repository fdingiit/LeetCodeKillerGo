package getIntersectionNode

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var len1, len2, diff int
	var p1, p2 *ListNode

	p1, p2 = headA, headB

	for p1 != nil {
		p1 = p1.Next
		len1++
	}

	for p2 != nil {
		p2 = p2.Next
		len2++
	}

	diff = len1 - len2
	p1, p2 = headA, headB

	if diff < 0 {
		for diff != 0 {
			p2 = p2.Next
			diff++
		}
	} else if diff > 0 {
		for diff != 0 {
			p1 = p1.Next
			diff--
		}
	}

	for p1 != nil && p2 != nil {
		if p1 == p2 {
			return p1
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	return nil
}
