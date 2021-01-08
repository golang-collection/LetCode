package code

/**
* @Author: super
* @Date: 2021-01-08 09:51
* @Description:
**/

func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int)
	left, right, res := 0, 0, 0
	for right < len(s) {
		c := s[right]
		right++
		window[c]++
		for window[c] > 1 {
			d := s[left]
			left++
			window[d]--
		}
		res = max(res, right-left)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
