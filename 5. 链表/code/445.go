package code

/**
* @Author: super
* @Date: 2021-01-08 21:44
* @Description: todo:重新做一下
**/

func addTwoNumbers(h1 *ListNode, h2 *ListNode) *ListNode {
	tmp1 := make([]int, 0, 10)
	tmp2 := make([]int, 0, 10)
	for ; h1 != nil; h1 = h1.Next {
		tmp1 = append(tmp1, h1.Val)
	}
	for ; h2 != nil; h2 = h2.Next {
		tmp2 = append(tmp2, h2.Val)
	}
	lenth1, lenth2 := len(tmp1), len(tmp2)
	var lenth int
	if lenth1 > lenth2 {
		tmp := make([]int, lenth1-lenth2)
		tmp2 = append(tmp, tmp2...)
		lenth = lenth1
	} else {
		tmp := make([]int, lenth2-lenth1)
		tmp1 = append(tmp, tmp1...)
		lenth = lenth2
	}
	c, sum := 0, 0
	res := make([]int, lenth)
	for i := lenth - 1; i >= 0; i-- {
		sum = tmp1[i] + tmp2[i] + c
		res[i] = sum % 10
		c = sum / 10
	}
	if c > 0 {
		res = append([]int{c}, res...)
	}
	return CreateNode(res)
}

func CreateNode(nums []int) *ListNode {
	lenth := len(nums)
	if lenth == 0 {
		return nil
	}
	res := &ListNode{
		Val:  nums[0],
		Next: nil,
	}
	next := CreateNode(nums[1:])
	res.Next = next
	return res
}