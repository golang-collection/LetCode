package code

/**
* @Author: super
* @Date: 2021-01-05 20:15
* @Description: 双指针
**/

func twoSum(numbers []int, target int) []int {
	if numbers == nil || len(numbers) == 0 {
		return nil
	}
	low, high := 0, len(numbers)-1
	for low < high {
		sum := numbers[low] + numbers[high]
		if sum == target {
			return []int{low+1, high+1}
		} else if sum > target {
			high--
		} else {
			low++
		}
	}
	return nil
}
