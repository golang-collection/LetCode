package main

/**
* @Author: super
* @Date: 2020-12-11 22:33
* @Description: https://leetcode-cn.com/problems/n-queens/
n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
**/

var res = make([][]string, 0)

func solveNQueens(n int) [][]string {
	res = make([][]string, 0)
	board := make([]string, n)
	for i := 0; i<n;i++{
		for j := 0; j<n; j++{
			board[i] += "."
		}
	}
	backtrack(board, 0)
	return res
}

func backtrack(board []string, row int) {
	// 触发结束条件
	if row == len(board) {
		res = append(res, board)
		return
	}

	n := len(board[row])
	for col := 0; col < n; col++ {
		// 排除不合法选择
		if !isValid(board, row, col){
			continue
		}
		// 做选择
		board[row][col] = 'Q'
		// 进入下一行决策
		backtrack(board, row + 1)
		// 撤销选择
		board[row][col] = '.'
	}
}

func isValid(board []string, row int, col int) bool {
	n := len(board)
	// 检查列是否有皇后互相冲突
	for i:= 0; i < n; i++ {
		if board[i][col] == 'Q'{
			return false
		}
	}
	// 检查右上方是否有皇后互相冲突
	for i := row - 1, j := col + 1; i >= 0 && j < n; i--, j++ {
		if board[i][j] == 'Q'{
			return false
		}
	}
	// 检查左上方是否有皇后互相冲突
	for i := row - 1, j = col - 1; i >= 0 && j >= 0; i--, j-- {
		if board[i][j] == 'Q'{
			return false
		}
	}
	return true
}