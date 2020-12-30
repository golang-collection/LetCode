package code

/**
* @Author: super
* @Date: 2020-12-30 16:16
* @Description:
**/

func mergeKLists(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists) - 1)
}

func merge(lists []*ListNode, l int, r int) *ListNode{
	if l == r{
		return lists[l]
	}
	if l > r{
		return nil
	}
	mid := (l + r) >> 1
	return mergeTwoLists(merge(lists, l, mid), merge(lists, mid + 1, r))
}