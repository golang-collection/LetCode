package code

/**
* @Author: super
* @Date: 2020-12-10 21:20
* @Description: https://leetcode-cn.com/problems/fibonacci-number/
  斐波那契数，通常用 F(n) 表示，形成的序列称为斐波那契数列。该数列由 0 和 1 开始，
  后面的每一项数字都是前面两项数字的和。也就是：
F(0) = 0,   F(1) = 1
F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
  给定 N，计算FN)。
**/

func fib(N int) int {
	if N < 1 {
		return 0
	} else if N == 2 || N == 1 {
		return 1
	}
	pre, curr := 1, 1
	for i := 1; i <= N; i++ {
		sum := pre + curr
		pre = curr
		curr = sum
	}
	return curr
}
