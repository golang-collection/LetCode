package code

/**
* @Author: super
* @Date: 2021-03-12 10:39
* @Description:
**/

func findRepeatNumber(nums []int) int {
	if nums == nil || len(nums) == 0{
		return -1
	}
	bitmap := [100000]int{}
	for i := 0; i<len(nums); i++{
		if bitmap[nums[i]] != 0{
			return nums[i]
		}
		bitmap[nums[i]]++
	}
	return -1
}