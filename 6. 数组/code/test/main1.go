package main

import (
	"fmt"
)

/**
* @Author: super
* @Date: 2021-03-21 20:18
* @Description:
**/

func main() {
	n := 0
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var ans int64
		fmt.Scan(&ans)
		fmt.Println(number(ans))
	}
}

var hmap  = make(map[int64]int64)

func number(n int64) int64 {
	if v, ok := hmap[n]; ok{
		return v
	}
	if n == 0 {
		return 0
	}
	if n%2 == 0 {
		a := min(number(n/2), number(n - 1)) + 1
		hmap[n] = a
		return a
	}
	if n%3 == 0 {
		a := min(number(n/3), number(n - 1)) + 1
		hmap[n] = a
		return min(number(n/3), number(n - 1)) + 1
	}
	return number(n-1) + 1
}

func min(a, b int64) int64 {
	if a > b {
		return b
	}
	return a
}
