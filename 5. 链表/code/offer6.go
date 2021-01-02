package code

/**
* @Author: super
* @Date: 2021-01-02 20:12
* @Description:
**/

func reversePrint(head *ListNode) []int {
	result := make([]int, 0)
	p := head
	for p != nil{
		result = append([]int{p.Val}, result...)
		p = p.Next
	}
	return result
}