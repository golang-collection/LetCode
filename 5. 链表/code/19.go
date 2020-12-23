package code

/**
* @Author: super
* @Date: 2020-12-22 21:34
* @Description: 没有头节点就先构造头节点在操作更简单
**/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	maskHead := &ListNode{0, head}
	fast, low := head, maskHead
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for ; fast != nil; fast = fast.Next {
		low = low.Next
	}
	low.Next = low.Next.Next
	return maskHead.Next
}
