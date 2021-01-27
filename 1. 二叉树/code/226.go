package code

/**
* @Author: super
* @Date: 2021-01-27 21:01
* @Description:
**/

func invertTree(root *TreeNode) *TreeNode {
	inverse(root)
	return root
}

func inverse(root *TreeNode){
	if root == nil || (root.Left == nil && root.Right == nil){
		return
	}
	temp := root.Left
	root.Left = root.Right
	root.Right = temp
	inverse(root.Left)
	inverse(root.Right)
}