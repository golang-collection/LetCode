package main

import (
	"fmt"
	"strconv"
)

func numberOf0(n int) int {
	count := 0
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	str := strconv.Itoa(result)
	for i := len(str) - 1; i > 0; i-- {
		if str[i] == '0' {
			count++
		} else {
			break
		}
	}
	return count
}

func main() {
	n := 0
	fmt.Scanf("%d", &n)
	fmt.Println(numberOf0(n))
}
