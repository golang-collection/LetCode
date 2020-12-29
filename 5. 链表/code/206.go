package code

/**
* @Author: super
* @Date: 2020-12-29 21:25
* @Description: 反转链表
**/

// 递归解法
func reverseList(head *ListNode) *ListNode {
	if head == nil{
		return head
	}
	if head.Next == nil {
		return head
	}
	last := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

//非递归解法
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	h := &ListNode{
		Next:head,
	}
	p := head.Next
	q := head.Next
	head.Next = nil
	for q != nil{
		q = p.Next
		p.Next = head
		h.Next = p
		head = p
		p = q
	}
	return h.Next
}