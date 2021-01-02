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
	if head == nil{
		return head
	}

}
