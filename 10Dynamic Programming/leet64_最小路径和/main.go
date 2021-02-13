package main

import "fmt"

/*
给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
说明：每次只能向下或者向右移动一步。

示例 1：
输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
输出：7
解释：因为路径 1→3→1→1→1 的总和最小。
示例 2：

输入：grid = [[1,2,3],[4,5,6]]
输出：12

提示：

m == grid.length
n == grid[i].length
1 <= m, n <= 200
0 <= grid[i][j] <= 100

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
动态规划:
	1.定义数组元素的含义
		设置一个数组dp[i][j]存放在(i,j)位置的最小和，因为网格式m x n的，数组下标从0开始，所以最小和为dp[m-1][n-1]
	2.找出数组元素之间的关系式
		导致(i,j)位置的最小和是(i,j-1)位置的最小和与(i-1,j)位置的最小和比较后加上(i,j)位置的值grid[i][j]
		dp[i][j] = min{dp[i][j-1],dp[i-1][j]} + grid[i][j]
	3.找初始值
		很容易得知，从当i,j有一方为0时，等式不成立，此时dp[0][j] = sum{grid[0][j]},dp[i][0]=sum{grid[i][0]}
	即表示为dp[0][j] = dp[0][j-1] + grid[0][j]   dp[i][0] = dp[i-1][0] + grid[i][0]
*/

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	if m == 0 || n == 0 {
		panic("error array")
	}
	//fmt.Println(m, n)

	//1.
	dp := make([][]int, m) //构建二维切片的方法有些特殊
	for i := range dp {
		dp[i] = make([]int, n)
	}

	//3.
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	//fmt.Println(dp)

	//2.
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if dp[i][j-1] <= dp[i-1][j] {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			}
		}
	}

	return dp[m-1][n-1]
}

func main() {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	pathMinSum := minPathSum(grid)
	fmt.Println(pathMinSum)

}
