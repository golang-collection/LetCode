package code

/**
* @Author: super
* @Date: 2021-01-02 20:18
* @Description: 编写代码，移除未排序链表中的重复节点。保留最开始出现的节点。

**/

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	occurred := map[int]bool{head.Val: true}
	pos := head
	for pos.Next != nil {
		cur := pos.Next
		if !occurred[cur.Val] {
			occurred[cur.Val] = true
			pos = pos.Next
		} else {
			pos.Next = pos.Next.Next
		}
	}
	pos.Next = nil
	return head
}