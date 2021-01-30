package code

/**
* @Author: super
* @Date: 2021-01-30 21:38
* @Description: 二叉排序树中插入节点
**/

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil{
		return &TreeNode{Val:val}
	}
	if root.Val < val{
		root.Right = insertIntoBST(root.Right, val)
	}else{
		root.Left = insertIntoBST(root.Left, val)
	}
	return root
}