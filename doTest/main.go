package main

import (
	"fmt"
	"strings"
)

/**
* @Author: super
* @Date: 2021-01-30 21:18
* @Description:
**/

var dp [100010][2]int
var v [][]int
var add [100010][2][]int

func dfs(node,pre int){
	for i := 0; i<len(v[node]);i++{
		temp := v[node][i]
		if temp!=pre{
			dfs(temp, node)
			dp[node][0] += max(dp[temp][0], dp[temp][1])
			dp[node][1] += dp[temp][0]
		}
	}
}

func max(a ,b int)int {
	if a > b{return a}
	return b
}

func main() {
	n,a,b := 0,0,0
	fmt.Scanf("%d", &n)
	for i := 1; i<=n;i++{
		fmt.Scanf("%d", &dp[i][1])
	}
	v := make([][]int, n+1)
	for i:= 1; i<=n-1;i++{
		fmt.Scanf("%d%d", &a,&b)
	}
}