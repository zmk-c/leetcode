package main

import (
	"fmt"
	"strings"
)

/*
参数说明：
row int	行数
n	皇后个数
chessboard	棋盘
result 结果集
*/
func backtrack(row, n int, chessboard [][]string, result *[][]string) {
	//终止条件  当行数row递归到棋盘最底层
	if row == n {
		tmp := make([]string, n)
		for i := 0; i < n; i++ {
			tmp[i] = strings.Join(chessboard[i], "") //将chessboard降维成一维数组 存入tmp中
		}
		*result = append(*result, tmp)
		return
	}

	//单层for循环搜索逻辑
	for col := 0; col < n; col++ { //列数每次从最开始遍历
		if isValid(row, col, n, chessboard) {
			chessboard[row][col] = "Q" //将合法位置上放入皇后
			backtrack(row+1, n, chessboard, result)
			chessboard[row][col] = "." //回溯
		}
	}
}

func isValid(row int, col int, n int, chessboard [][]string) bool {
	for i := 0; i < row; i++ {
		for j := 0; j < n; j++ {
			if chessboard[i][j] == "Q" && (i+j == row+col || i-j == row-col || j == col) { //剪枝函数
				return false
			}
		}
	}
	return true
}

func solveNQueens(n int) [][]string {
	chessboard := make([][]string, n)
	for i := 0; i < n; i++ {
		chessboard[i] = make([]string, n)
		for j := 0; j < n; j++ {
			chessboard[i][j] = "."
		}
	}
	result := make([][]string, 0)
	backtrack(0, n, chessboard, &result)
	return result
}

func main() {
	n := 4 //n 表示皇后的个数
	queens := solveNQueens(n)
	fmt.Printf("%v皇后的解决方案：\n%v", n, queens)
}
