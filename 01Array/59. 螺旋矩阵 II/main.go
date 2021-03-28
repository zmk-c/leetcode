package main

/*
这题主要考咱们的是模拟能力和对代码的掌握能力，特别是对边界的考虑。

我们可以将填充矩阵元素分为以下四个过程:

- 填充上行从左到右
- 填充右列从上到下
- 填充下行从右到左
- 填充左列从下到上

这里可以看到需要引入left,right,top,bottom四个变量作为边界的控制，
因为矩阵从0开始，所以四个变量分别初始化为left=0 right=n-1 top=0 bottom=n-1 （其中n为给定正整数）
接下里就是按照填充顺序进行模拟。
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
