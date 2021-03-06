package main

import "fmt"

func backtrack(board [][]byte) bool {
	for i := 0; i < len(board); i++ { //遍历行
		for j := 0; j < len(board[0]); j++ { //遍历列
			if board[i][j] != '.' {
				continue //说明该位置不是空白格  有数字了
			}
			//接下来判断(i,j)这个位置放入k是否合适
			for k := '1'; k <= '9'; k++ {
				if isValid(i, j, byte(k), board) {
					//位置合法放置
					board[i][j] = byte(k)
					if backtrack(board) { //如果找到合适的一组  就立刻返回
						return true
					}
					board[i][j] = '.' //回溯
				}
			}
			return false //如果9个数都试完了  证明无解
		}
	}
	return true // 遍历完没有返回false，说明找到了合适棋盘位置了
}

/*判断棋盘是否合法有如下三个维度：
同行是否重复
同列是否重复
9宫格里是否重复
*/
func isValid(row int, col int, val byte, board [][]byte) bool {
	for i := 0; i < 9; i++ {
		//判断行里是否重复
		if board[row][i] == val {
			return false
		}
		//判断列里是否重复
		if board[i][col] == val {
			return false
		}

		// 判断9宫格是否存在重复
		if board[(row/3)*3+i/3][(col/3)*3+i%3] == val {
			return false
		}
	}

	return true
}

func solveSudoku(board [][]byte) {
	backtrack(board)
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(board)
	for i := 0; i < len(board); i++ {
		fmt.Printf("%c \n", board[i])
	}
}
