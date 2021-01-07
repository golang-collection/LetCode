package code

import "math"

/**
* @Author: super
* @Date: 2021-01-05 20:32
* @Description: 滑动窗口
**/

func minWindow(s string, t string) string {
	need, window := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	left, right, valid := 0, 0, 0
	start, length := 0, math.MaxInt64
	for right < len(s){
		c := s[right]
		right++
		if _, ok := need[c]; ok{
			window[c]++
			if window[c] == need[c]{
				valid++
			}
		}
		for len(need) == valid{
			if right - left < length{
				start = left
				length = right - left
			}
			d := s[left]
			left++
			if _, ok := need[d]; ok{
				if window[d] == need[d]{
					valid--
				}
				window[d]--
			}
		}
	}
	if length == math.MaxInt64{
		return ""
	}
	return s[start: start + length]
}
