package code

/**
* @Author: super
* @Date: 2021-01-02 20:25
* @Description:
**/

func kthToLast(head *ListNode, k int) int {
	if head == nil{
		return 0
	}
	slow, fast := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for ; fast != nil; fast = fast.Next {
		slow = slow.Next
	}
	return slow.Val
}