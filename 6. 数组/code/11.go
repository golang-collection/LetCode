package code

/**
* @Author: super
* @Date: 2021-01-06 21:03
* @Description:
**/

func maxArea(height []int) int {
	if height == nil || len(height) == 0 {
		return 0
	}
	maxArea, left, right := 0, 0, len(height)-1
	for left < right {
		area := (right - left) * min(height[left], height[right])
		if area > maxArea {
			maxArea = area
		}
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return maxArea
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}