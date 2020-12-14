package main

import "fmt"

/**
* @Author: super
* @Date: 2020-12-11 11:31
* @Description: https://leetcode-cn.com/problems/permutations/
  给定一个 没有重复 数字的序列，返回其所有可能的全排列。
**/

var res = make([][]int, 0)

func permute(nums []int) [][]int {
	res = make([][]int, 0)
	track := make([]int, 0)
	backtrack(nums, track)
	return res
}

func backtrack(nums []int, track []int) {
	if len(nums) == len(track) {
		// 不复制会导致track继续被更改导致结果错误
		temp := make([]int, len(nums))
		copy(temp, track)
		res = append(res, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if ok := contains(track, nums[i]); ok {
			continue
		}
		track = append(track, nums[i])
		backtrack(nums, track)
		track = track[:len(track)-1]
	}
}

func contains(track []int, num int) bool {
	if track == nil || len(track) == 0 {
		return false
	}
	for i := 0; i < len(track); i++ {
		if track[i] == num {
			return true
		}
	}
	return false
}

func main() {
	arr := []int{5, 4, 6, 2}
	result := permute(arr)
	fmt.Print(result)
}
