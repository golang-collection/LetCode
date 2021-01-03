package code

/**
* @Author: super
* @Date: 2021-01-03 16:25
* @Description:
**/

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	newHead := &ListNode{
		Next:nil,
	}
	p, q := head, head.Next
	for q != nil{
		h := newHead
		for h.Next != nil || h.Next.Val < p.Val{
			h = h.Next
		}
		p.Next = h.Next
		h.Next = p
		p = q
		q = q.Next
	}
	return newHead
}