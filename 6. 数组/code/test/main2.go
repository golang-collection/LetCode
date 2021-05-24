package main

import (
	"fmt"
	"math"
)

/**
* @Author: super
* @Date: 2021-03-22 19:04
* @Description:
**/

func sum(n int, w []int, target int) {
	l := int(math.Pow(float64(2), float64(n+1)))
	tree := make([]int, l)
	tree[1] = 0
	i := 0
	level := 4
	for id := 2; id < l; id++ {
		if id >= level {
			i += 1
			level *= 2
		}
		if id%2 == 0 {
			tree[id] += w[i] + tree[id/2]
		} else {
			tree[id] += tree[id/2]
		}
		if tree[id] == target {
			fmt.Println("YES")
			return
		}
	}
	fmt.Println("NO")
}

func main() {
	n := 0
	target := 0
	fmt.Scanf("%d %d", &n, &target)
	for {
		w := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &w[i])
		}
		sum(n, w, target)
		t, err := fmt.Scanf("%d %d", &n, &target)
		if t == 0 || err != nil{
			break
		}
	}
}
