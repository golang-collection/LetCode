package code

/**
* @Author: super
* @Date: 2020-12-24 20:55
* @Description:
**/

func search(nums []int, target int) int {
	if nums == nil || len(nums) == 0{
		return -1
	}
	low, high := 0, len(nums) - 1
	for low <= high{
		// 如果low和right太大下面这种写法会导致溢出
		//mid := (low + high) / 2
		mid := low + (high - low) / 2
		if nums[mid] == target{
			return mid
		}else if nums[mid] > target{
			high = mid - 1
		}else if nums[mid] < target{
			low = mid + 1
		}
	}
	return -1
}