package code

/**
* @Author: super
* @Date: 2021-01-03 16:02
* @Description:
**/

func reorderList(head *ListNode) {
	if head != nil {
		slow, fast := head, head
		for fast != nil && fast.Next != nil {
			fast = fast.Next.Next
			slow = slow.Next
		}
		temp := slow
		slow = slow.Next
		temp.Next = nil
		slow = reverse(slow, nil)
		p := slow
		fast = head
		for p != nil {
			p = p.Next
			slow.Next = fast.Next
			fast.Next = slow
			fast = slow.Next
			slow = p
		}
	}
}
