package code

/**
* @Author: super
* @Date: 2021-01-08 09:44
* @Description:
**/

func findAnagrams(s string, p string) []int {
	need := make(map[byte]int)
	window := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}

	left, right, valid := 0, 0, 0
	result := make([]int, 0)
	for right < len(s) {
		c := s[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		for right-left >= len(p) {
			if valid == len(need) {
				result = append(result, left)
			}
			d := s[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return result
}
