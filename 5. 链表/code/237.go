package code

/**
* @Author: super
* @Date: 2021-01-02 20:07
* @Description:
**/

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}