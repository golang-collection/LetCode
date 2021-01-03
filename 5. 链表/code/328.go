package code

/**
* @Author: super
* @Date: 2021-01-03 16:56
* @Description:
**/

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	pre, slow, fast := head, head.Next, head.Next.Next
	for fast != nil{
		slow.Next = fast.Next
		fast.Next = pre.Next
		pre.Next = fast
		pre = fast
		slow = slow.Next
		if slow == nil{
			break
		}
		fast = slow.Next
	}
	return head
}