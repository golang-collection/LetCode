package code

/**
* @Author: super
* @Date: 2021-01-13 07:22
* @Description:
**/

var result []int

func inorderTraversal(root *TreeNode) []int {
	result = make([]int, 0)
	inorder(root)
	return result
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	result = append(result, root.Val)
	inorder(root.Right)
}

func inorderTraversal(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return
}