package code

/**
* @Author: super
* @Date: 2021-01-12 23:20
* @Description:
**/

func isSymmetric(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil{
		return true
	}
	if root.Left == nil || root.Right == nil{
		return false
	}
	if root.Left.Val != root.Right.Val{
		return true
	}
	return isSymmetric(root.Left) && isSymmetric(root.Right)
}