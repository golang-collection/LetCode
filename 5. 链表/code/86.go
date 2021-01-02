package code

/**
* @Author: super
* @Date: 2021-01-02 21:48
* @Description:
**/

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	newHead := &ListNode{
		Next:nil,
	}
	h, tail := newHead, newHead
	p, q := head, head.Next
	for p != nil{
		if p.Val >= x{
			p.Next = tail.Next
			tail.Next = p
			tail = p
		}else{
			if h.Next == nil{
				tail = p
			}
			p.Next = h.Next
			h.Next = p
			h = p
		}
		p = q
		if q != nil{
			q = q.Next
		}
	}
	return newHead.Next
}