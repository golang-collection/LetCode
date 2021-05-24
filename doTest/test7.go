package main

import (
	"fmt"
	"io"
)

func union(n, k int, power, agree []int) int {
	max := 0
	if k >= n{
		for _, i := range power{
			max += i
		}
		return max
	}
	for i := 0; i < n - k ;i++{
		temp := 0
		for j := 0; j < k; j++ {
			temp += power[i + j]
		}
		low := i - 1
		high := i + k
		for low > 0 && agree[low] == 1{
			temp += power[low]
			low--
		}
		for high < n && agree[high] == 1{
			temp += power[high]
			high++
		}
		if temp > max{
			max = temp
		}
	}
	return max
}

func main() {
	var n, k int
	var power, agree []int
	for {
		if _, err := fmt.Scan(&n,&k); err != io.EOF {
			power = make([]int, n)
			agree = make([]int, n)
			for i := 0; i<n; i++{
				fmt.Scan(&power[i])
			}
			for i := 0; i<n; i++{
				fmt.Scan(&agree[i])
			}
			fmt.Println(union(n,k, power, agree))
		} else {
			break
		}
	}
}