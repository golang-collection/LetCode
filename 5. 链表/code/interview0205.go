package code

/**
* @Author: super
* @Date: 2021-01-02 20:32
* @Description:
**/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}else if l1 == nil{
		return l2
	}else if l2 == nil{
		return l1
	}
	head := &ListNode{
		Val:-1,
		Next:nil,
	}
	tail := head

	flag := 0
	for l1 != nil || l2 != nil{
		val1, val2 := 0, 0
		if l1 != nil{
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil{
			val2 = l2.Val
			l2 = l2.Next
		}

		sum := val1 + val2 + flag
		if sum >= 10 {
			flag = 1
			sum -= 10
		}else{
			flag = 0
		}
		tail = listAdd(tail, sum)
	}
	if flag == 1{
		tail = listAdd(tail, 1)
	}
	return head.Next
}