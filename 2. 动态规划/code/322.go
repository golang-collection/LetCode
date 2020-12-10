package code

import "math"

/**
* @Author: super
* @Date: 2020-12-10 21:34
* @Description: https://leetcode-cn.com/problems/coin-change/
给定不同面额的硬币 coins 和一个总金额 amount。
编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
你可以认为每种硬币的数量是无限的。
**/
//dp数组解法
func coinChange1(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i<len(dp); i++{
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 0; i<len(dp); i++{
		for _, coin := range coins{
			if i - coin < 0{
				continue
			}
			dp[i] = min(dp[i],  1+ dp[i - coin])
		}
	}
	if dp[amount] == amount + 1{
		return -1
	}
	return dp[amount]
}

func coinChange(coins []int, amount int) int {
	memo := make(map[int]int)
    return dp(coins, memo, amount)
}

func dp(coins []int, memo map[int]int, n int) int {
	if n == 0{
		return 0
	}
	if n < 0{
		return -1
	}
	if v, ok := memo[n]; ok{
		return v
	}
	res := math.MaxInt8
	for _, coin := range coins{
		subProblem := dp(coins, memo, n-coin)
		if subProblem == -1{
			continue
		}
		res = min(res, 1 + subProblem)
	}
	if res != math.MaxInt8{
		memo[n] = res
	}else{
		memo[n] = -1
	}
	return memo[n]
}

func min(a, b int)int{
	if a < b {
		return a
	}
	return b
}