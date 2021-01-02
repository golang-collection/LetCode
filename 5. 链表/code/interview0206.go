package code

/**
* @Author: super
* @Date: 2021-01-02 20:29
* @Description:
**/

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil{
		return true
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil{
		fast = fast.Next.Next
		slow = slow.Next
	}
	if fast != nil{
		slow = slow.Next
	}
	slow = reverse(slow, nil)
	fast = head
	for slow != nil{
		if slow.Val != fast.Val{
			return false
		}
		slow = slow.Next
		fast = fast.Next
	}
	return true
}