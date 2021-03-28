package main

import "fmt"

//在leet77_组合问题的求解上加上了总和这个限制条件
/*
参数解释：
	k int 题目需要的每个集合中有k个数
	n int 题目条件，k个数相加的和
	sum int 在寻找过程中的累加和 动态
	startIndex int 开始遍历的下标
	path []int 收集符合条件的元素
	result *[][]int 结果数组
*/
func backtrack(k, n, sum, startIndex int, path []int, result *[][]int) {
	if sum > n { //剪枝函数  如果len(path)==k 而 sum != n 直接返回
		return
	}

	//终止条件 path结果集中的个数满足题目要求的k个数退出
	if k == len(path) {
		if n == sum { //限制条件，当前path集合元素的累加和sum是否等于题目给定的和n
			//将满足条件的一个解path加入最终结果集result中
			tmp := make([]int, k)
			copy(tmp, path)
			*result = append(*result, tmp)

			return
		}
	}

	//单个循环判断
	//剪枝函数 和leet77_组合问题一样
	for i := startIndex; i <= 9-(k-len(path))+1; i++ {
		sum += i
		path = append(path, i)
		//进行下一层的递归
		backtrack(k, n, sum, i+1, path, result)
		//回溯
		sum -= i
		path = path[:len(path)-1]
	}
}

func combinationSum3(k int, n int) [][]int {
	path := make([]int, 0)
	result := make([][]int, 0)
	backtrack(k, n, 0, 1, path, &result)
	return result
}

func main() {
	result := combinationSum3(3, 9)
	fmt.Println(result)
}
