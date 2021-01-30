package code

/**
* @Author: super
* @Date: 2021-01-30 21:47
* @Description:
**/

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil{
		return nil
	}
	if root.Val == key{
		if root.Left == nil{
			return root.Right
		}else if root.Right == nil{
			return root.Left
		}
		minNode := getMin(root.Right)
		root.Val = minNode.Val
		root.Right = deleteNode(root.Right, minNode.Val)
	}else if root.Val > key{
		root.Left = deleteNode(root.Left, key)
	}else if root.Val < key{
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

func getMin(root *TreeNode) *TreeNode {
	for root.Left != nil{
		root = root.Left
	}
	return root
}