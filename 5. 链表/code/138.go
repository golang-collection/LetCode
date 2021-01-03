package code

/**
* @Author: super
* @Date: 2021-01-02 22:07
* @Description:
**/

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	p := head
	for p != nil {
		temp := &Node{
			Val:  p.Val,
			Next: p.Next,
		}
		p.Next = temp
		p = temp.Next
	}
	p, q, newHead := head, head.Next, head.Next
	for q != nil {
		if p.Random == nil{
			q.Random = nil
		}else{
			q.Random = p.Random.Next
		}
		if q.Next == nil{
			break
		}
		p = q.Next
		q = q.Next.Next
	}
	p, q = head, head.Next
	for q.Next != nil{
		p.Next = q.Next
		q.Next = q.Next.Next
		p = p.Next
		q = q.Next
	}
	p.Next = q.Next
	return newHead
}
