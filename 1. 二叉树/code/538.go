package code

/**
* @Author: super
* @Date: 2021-01-30 21:10
* @Description:
**/
var sum int

func convertBST(root *TreeNode) *TreeNode {
	traverse(root)
	return root
}

func traverse(root *TreeNode) {
	if root == nil{
		return
	}
	traverse(root.Right)
	sum += root.Val
	root.Val = sum
	traverse(root.Left)
}

func convertBST(root *TreeNode) *TreeNode {
	sums := 0
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode){
		if root==nil{
			return
		}
		dfs(root.Right)
		sums = sums+root.Val
		root.Val = sums
		dfs(root.Left)
		return
	}
	dfs(root)
	return root
}