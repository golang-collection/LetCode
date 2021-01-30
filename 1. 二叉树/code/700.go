package code

/**
* @Author: super
* @Date: 2021-01-30 21:35
* @Description: 二叉排序树中搜索节点
**/

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil{
		return nil
	}
	if root.Val == val{
		return root
	}else if root.Val < val{
		return searchBST(root.Right, val)
	}else{
		return searchBST(root.Left, val)
	}
}