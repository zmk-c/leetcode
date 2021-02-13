package main

import "fmt"

/*
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
问总共有多少条不同的路径？

示例 1：
输入：m = 3, n = 7
输出：28

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-paths
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
动态规划三步走：
	1.定义数组元素的含义
		这里我们定义dp[i][j]的含义为，机器人从左上角走到(i,j)这个位置共走了多少步，因为网格为m x n的，数组下标从0开始
	所有总的总路数为d[m-1][n-1]
	2.找出数组元素之间的关系式
		到达(i,j)这个位置有两种走法，即从(i,j-1)到和(i-1,j)到达，所以dp[i][j]=dp[i][j-1]+dp[i-1][j]
	3.找初始值
		很明显我们可以发现，当i或j有一方为0时，等式不成立，出现了负数，所以设置初始值计算出dp[0][0...n]<-1 //机器人一直往右走
	dp[0...m][0]<-1 //机器人一直向下走。
*/
func uniquePaths(m int, n int) int {
	//时间复杂度O(m*n)  每个网格都需要走

	//1
	dp := make([][]int, m) //二维切片的创造有点独特  注意一下
	for i := range dp {
		dp[i] = make([]int, n)
	}
	//fmt.Println(dp)

	//3
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	//2
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}

	return dp[m-1][n-1]
}

func main() {
	m, n := 3, 7
	sumPath := uniquePaths(m, n)
	fmt.Println(sumPath)
}
