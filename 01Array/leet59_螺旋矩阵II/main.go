package main

import "fmt"

/*
给定一个正整数 n，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。

示例:
输入: 3
输出:
[
 [ 1, 2, 3 ],
 [ 8, 9, 4 ],
 [ 7, 6, 5 ]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/spiral-matrix-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func generateMatrix(n int) [][]int {
	left, right, top, bottom := 0, n-1, 0, n-1
	//创建二维数组
	martix := make([][]int, n)
	for i := 0; i < n; i++ {
		martix[i] = make([]int, n)
	}
	num, target := 1, n*n
	/*0 1 2
	0[1,2,3
	1 8,9,4
	2 7,6,5]
	*/
	for num <= target {
		for i := left; i <= right; i++ { //从left到right
			martix[top][i] = num // (0,0)=>1  (0,1)=>2   (0,2)=>3     //(1,1)=>9
			num++
		}
		top++ //1

		for i := top; i <= bottom; i++ { //从top到bottom
			martix[i][right] = num //(1,2)=>4  (2,2)=>5
			num++
		}
		right-- //1

		for i := right; i >= left; i-- { //从right到left
			martix[bottom][i] = num //(2,1)=>6  (2,0)=>7
			num++
		}
		bottom-- //1

		for i := bottom; i >= top; i-- { //从bottom到top
			martix[i][left] = num //(1,0)=>8
			num++
		}
		left++ //1
	}
	return martix
}

func main() {
	matrix := generateMatrix(3)
	fmt.Println(matrix)
}
