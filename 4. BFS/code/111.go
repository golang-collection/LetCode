package code

/**
* @Author: super
* @Date: 2020-12-14 21:30
* @Description:
**/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil{
		return 0
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	depth := 1
	for len(q) != 0{
		sz := len(q)
		for i := 0; i<sz; i++{
			cur := q[0]
			q = q[1:]
			if cur.Left == nil && cur.Right == nil{
				return depth
			}
			if cur.Left != nil{
				q = append(q, cur.Left)
			}
			if cur.Right != nil{
				q = append(q, cur.Right)
			}
		}
		depth++
	}
	return depth
}
