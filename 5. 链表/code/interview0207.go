package code

/**
* @Author: super
* @Date: 2021-01-02 20:30
* @Description:
**/

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil{
		return nil
	}
	pA, pB := headA, headB
	for pA != pB{
		if pA == nil{
			pA = headB
		}else{
			pA = pA.Next
		}
		if pB == nil{
			pB = headA
		}else{
			pB = pB.Next
		}
	}
	return pA
}