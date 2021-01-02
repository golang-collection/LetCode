package code

/**
* @Author: super
* @Date: 2021-01-02 20:08
* @Description:
**/

func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil{
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}