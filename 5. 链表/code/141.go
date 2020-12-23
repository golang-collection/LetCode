package code

/**
* @Author: super
* @Date: 2020-12-22 21:54
* @Description:
**/

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil{
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow{
			return true
		}
	}
	return false
}
