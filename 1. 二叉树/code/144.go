package code

/**
* @Author: super
* @Date: 2021-01-13 07:30
* @Description:
**/
var result []int

func preorderTraversal(root *TreeNode) []int {
	result = make([]int, 0)
	preorder(root)
	return result
}

func preorder(root *TreeNode) {
	if root == nil {
		return
	}
	result = append(result, root.Val)
	preorder(root.Left)
	preorder(root.Right)
}