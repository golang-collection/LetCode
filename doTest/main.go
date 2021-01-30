package main

import "fmt"

/**
* @Author: super
* @Date: 2021-01-30 21:18
* @Description:
**/

func main() {
	a := 1
	b := 3
	a, b = a+b, b+a
	fmt.Println(a, b)
}