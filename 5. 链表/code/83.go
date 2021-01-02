package code

/**
* @Author: super
* @Date: 2021-01-02 16:43
* @Description:
**/

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	p, q := head, head.Next
	for q != nil{
		if q.Val == p.Val{
			p.Next = q.Next
			q.Next = nil
			q = p.Next
		}else{
			p = p.Next
			q = q.Next
		}
	}
	return head
}