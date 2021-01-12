package code

/**
* @Author: super
* @Date: 2021-01-13 07:32
* @Description:
**/

var result []int

func postorderTraversal(root *TreeNode) []int {
	result = make([]int, 0)
	postorder(root)
	return result
}

func postorder(root *TreeNode) {
	if root == nil {
		return
	}
	postorder(root.Left)
	postorder(root.Right)
	result = append(result, root.Val)
}