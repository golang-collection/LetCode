package code

import "math"

/**
* @Author: super
* @Date: 2021-01-28 21:34
* @Description:
**/

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return build(nums, 0, len(nums)-1)
}

func build(nums []int, low int, high int) *TreeNode {
	if low > high {
		return nil
	}
	index, maxValue := -1, math.MinInt64
	for i := low; i <= high; i++ {
		if maxValue < nums[i] {
			index = i
			maxValue = nums[i]
		}
	}

	root := &TreeNode{
		Val: maxValue,
	}
	root.Left = build(nums, low, index-1)
	root.Right = build(nums, index+1, high)
	return root
}
