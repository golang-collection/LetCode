package code

/**
* @Author: super
* @Date: 2021-01-05 20:20
* @Description: 双指针
**/

func reverseString(s []byte)  {
	if s == nil || len(s) == 0{
		return
	}
	low, high := 0, len(s) - 1
	for low < high{
		temp := s[low]
		s[low] = s[high]
		s[high] = temp
		low++
		high--
	}
}