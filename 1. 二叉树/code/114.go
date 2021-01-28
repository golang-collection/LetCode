package code

/**
* @Author: super
* @Date: 2021-01-25 22:12
* @Description:
**/

func flatten(root *TreeNode)  {
	curr := root
	for curr != nil {
		if curr.Left != nil {
			next := curr.Left
			predecessor := next
			for predecessor.Right != nil {
				predecessor = predecessor.Right
			}
			predecessor.Right = curr.Right
			curr.Left, curr.Right = nil, next
		}
		curr = curr.Right
	}
}

func flatten(root *TreeNode){
	if root == nil{
		return
	}
	flatten(root.Left)
	flatten(root.Right)

	left := root.Left
	right := root.Right

	root.Left = nil
	root.Right = left

	p := root
	for p.Right != nil{
		p = p.Right
	}
	p.Right = right
}