package code

/**
* @Author: super
* @Date: 2021-01-02 20:48
* @Description:
**/

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0{
		return head
	}
	slow, fast := head, head
	for ;k != 0;k--{
		if fast == nil{
			fast = head
		}
		fast = fast.Next
	}
	if fast == nil{
		return head
	}
	for fast.Next != nil{
		slow = slow.Next
		fast = fast.Next
	}
	h := slow.Next
	slow.Next = fast.Next
	fast.Next = head
	return h
}