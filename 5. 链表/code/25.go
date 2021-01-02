package code

/**
* @Author: super
* @Date: 2020-12-30 17:48
* @Description:
**/

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 1{
		return head
	}
	a, b := head, head
	for i := 0; i<k;i++{
		if b == nil{
			return head
		}
		b = b.Next
	}
	newHead := reverse(a, b)
	a.Next = reverseKGroup(b ,k)
	return newHead

}

func reverse(a, b *ListNode) *ListNode{
	var pre *ListNode
	cur, nxt := a, a
	for cur != b{
		nxt = cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}