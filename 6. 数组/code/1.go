package code

import "fmt"

/**
* @Author: super
* @Date: 2021-03-09 17:00
* @Description:
**/

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}

func find(n int, sum int) [][]int{
	low, high := 1, n
	result := make([][]int, 0)
	for low < high{
		if low + high < sum{
			low++
		}else if low + high > sum{
			high--
		}else if low + high == sum{
			result = append(result, []int{low, high})
		}
	}
	return result
}

func main() {
	//a := 0
	//fmt.Scan(&a)
	//fmt.Printf("%d\n", a)
	fmt.Println(find(5,3))
}