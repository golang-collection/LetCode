package code

/**
* @Author: super
* @Date: 2020-12-22 21:58
* @Description:
**/

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil{
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow{
			break
		}
	}
	if fast == nil || fast.Next == nil{
		return nil
	}
	slow = head
	for slow != fast{
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}