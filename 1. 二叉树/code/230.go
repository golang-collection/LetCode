package code

/**
* @Author: super
* @Date: 2021-01-30 20:54
* @Description: 二叉搜索树中第k小的元素
**/

//想找到第k小的元素，或者说找到排名为k的元素，如果想达到对数级复杂度，关键也在于每个节点得知道他自己排第几。
//
//比如说你让我查找排名为k的元素，当前节点知道自己排名第m，那么我可以比较m和k的大小：
//
//1、如果m == k，显然就是找到了第k个元素，返回当前节点就行了。
//
//2、如果k < m，那说明排名第k的元素在左子树，所以可以去左子树搜索第k个元素。
//
//3、如果k > m，那说明排名第k的元素在右子树，所以可以去右子树搜索第k - m - 1个元素。

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	lsize := size(root.Left)
	if lsize == k-1 {
		return root.Val
	} else if lsize < k-1 {
		return kthSmallest(root.Right, k-lsize-1)
	} else {
		return kthSmallest(root.Left, k)
	}
}

func size(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return size(root.Left) + size(root.Right) + 1
}
