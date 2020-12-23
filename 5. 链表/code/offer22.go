package code

/**
* @Author: super
* @Date: 2020-12-23 08:12
* @Description:
**/

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for ; fast != nil; fast = fast.Next {
		slow = slow.Next
	}
	return slow
}
