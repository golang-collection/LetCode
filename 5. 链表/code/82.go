package code

/**
* @Author: super
* @Date: 2021-01-02 20:59
* @Description:
**/

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := &ListNode{
		Next: head,
	}
	p, q, r := newHead, head, head.Next
	dub := false
	for r != nil {
		if q.Val == r.Val {
			r = r.Next
			dub = true
		} else {
			if dub == true {
				p.Next = r
				q = r
				r = r.Next
				dub = false
			} else {
				p = p.Next
				q = q.Next
				r = r.Next
			}
		}
	}
	if dub == true{
		p.Next = nil
	}
	return newHead.Next
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Val: 0, Next: head}
	cur, next := dummy, head
	for next != nil && next.Next != nil {
		if next.Val != next.Next.Val {
			cur.Next = next
			cur = cur.Next
			next = cur.Next
			continue
		}
		for next != nil && next.Next != nil && next.Val == next.Next.Val {
			next = next.Next
		}
		cur.Next = next.Next
		next = next.Next
	}

	return dummy.Next
}
