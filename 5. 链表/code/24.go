package code

/**
* @Author: super
* @Date: 2021-01-02 20:35
* @Description:
**/

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	newHead := &ListNode{
		Next:head,
	}
	h, p ,q := newHead, head, head.Next
	for q != nil{
		h.Next = q
		q = q.Next
		h.Next.Next = p
		p.Next = q
		h = p
		p = q
		if q != nil{
			q = q.Next
		}
	}
	return newHead.Next
}