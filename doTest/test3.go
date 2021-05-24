package main

import "fmt"

func twoSum(nums []int, target int) bool {
	hashTable := map[int]int{}
	for i, x := range nums {
		if _, ok := hashTable[target-x]; ok {
			return true
		}
		hashTable[x] = i
	}
	return false
}

func main() {
	fmt.Println(twoSum([]int{2,4,6,3,7,4,10}, 10))
}