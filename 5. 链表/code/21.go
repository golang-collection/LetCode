package code

/**
* @Author: super
* @Date: 2020-12-30 16:15
* @Description:
**/

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}

	p := head
	for l1 != nil && l2 != nil{
		if l1.Val <= l2.Val{
			p.Next = l1
			l1 = l1.Next
		}else{
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	if l1 == nil{
		p.Next = l2
	}else{
		p.Next = l1
	}
	return head.Next
}