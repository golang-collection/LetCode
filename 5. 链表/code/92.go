package code

/**
* @Author: super
* @Date: 2020-12-29 21:18
* @Description: 反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
**/
var successor *ListNode

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if n == 1{
		successor = head.Next
		return head
	}
	last := reverseBetween(head.Next, n-1)

}