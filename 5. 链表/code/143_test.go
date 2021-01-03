package code

import (
	"fmt"
	"testing"
)

/**
* @Author: super
* @Date: 2021-01-03 16:17
* @Description:
**/

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestReorderList(t *testing.T) {
	n3 := &ListNode{
		Next: nil,
		Val:  4,
	}
	n2 := &ListNode{
		Next: n3,
		Val:  3,
	}
	n1 := &ListNode{
		Next: n2,
		Val:  2,
	}
	head := &ListNode{
		Next: n1,
		Val:  1,
	}
	printList(head)
	reorderList(head)
	printList(head)
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Println(head)
		head = head.Next
	}
}

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

func reverse(a, b *ListNode) *ListNode {
	var pre *ListNode
	cur, nxt := a, a
	for cur != b {
		nxt = cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}
