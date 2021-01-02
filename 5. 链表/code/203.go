package code

/**
* @Author: super
* @Date: 2021-01-02 18:21
* @Description:
**/

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil{
		return head
	}
	newHead := &ListNode{
		Next:head,
	}
	p, q := newHead, head
	for q != nil{
		if q.Val == val{
			p.Next = q.Next
			q.Next = nil
			q = p.Next
		}else{
			p = p.Next
			q = q.Next
		}
	}
	return newHead.Next
}