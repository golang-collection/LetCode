package code

/**
* @Author: super
* @Date: 2021-01-15 08:58
* @Description: 给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？
**/

func numTrees(n int) int {
	G := make([]int, n + 1)
	G[0], G[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			G[i] += G[j-1] * G[i-j]
		}
	}
	return G[n]
	14
	34
}