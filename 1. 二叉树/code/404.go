package code

/**
* @Author: super
* @Date: 2021-02-02 23:12
* @Description:
**/

func sumOfLeftLeaves(root *TreeNode) (ans int) {
	dfs(root, &ans)
	return
}

func dfs(node *TreeNode, ans *int) {
	if node == nil {
		return
	}
	// 判断一下是不是左叶子节点
	if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
		*ans += node.Left.Val
	}
	dfs(node.Left, ans)
	dfs(node.Right, ans)
}