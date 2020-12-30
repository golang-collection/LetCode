package code

/**
* @Author: super
* @Date: 2020-12-29 21:18
* @Description: 反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
**/
var successor *ListNode

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == 1{
		return reverseN(head, n)
	}
	head.Next = reverseBetween(head.Next, m-1, n-1)
	return head
}

// 反转以 head 为起点的 n 个节点，返回新的头结点
func reverseN(head *ListNode , n int) *ListNode {
	if n == 1 {
		successor = head.Next
		return head
	}
	// 以 head.next 为起点，需要反转前 n - 1 个节点
	last := reverseN(head.Next, n - 1)

	head.Next.Next = head
	// 让反转之后的 head 节点和后面的节点连起来
	head.Next = successor
	return last
}