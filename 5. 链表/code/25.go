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
	newHead := &ListNode{
		Next:head,
	}

}

func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		nex := p.Next
		p.Next = prev
		prev = p
		p = nex
	}
	return tail, head
}