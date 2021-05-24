package main

import "fmt"

/**
* @Author: super
* @Date: 2021-03-19 10:19
* @Description:
**/

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func solve(root *TreeNode, b []int) []*ListNode {
	// write code here
    result := make([]*ListNode, 0)
	for i := 0; i<len(b); i++{
		head := &ListNode{}
		look(head, root, b[i])
		result = append(result, head.Next)
	}
	fmt.Println(result)
	return result
}

func addNode(head *ListNode, val int){
	t := head
	for t.Next != nil{
		t = t.Next
	}
	node := &ListNode{
		Val:val,
	}
	t.Next = node
}

func look(head *ListNode, root *TreeNode, val int) {
	if root == nil{
		return
	}
	addNode(head, root.Val)
	if root.Val == val{
		return
	}
	look(head, root.Left, val)
	look(head, root.Right, val)
}

func main() {

}