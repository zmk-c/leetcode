package main

import (
	"fmt"
	"sort"
)

//本题与leet77和leet216的区别：本题没有组合数量的要求，即没有递归层数的要求，但是有总和的限制，所以间接上也是有限制的
/*
参数解释：
candidates []int 题目所给的数组
target int 题目所给的目标和
sum int path路径上的累加和 需要与target进行判断
startIndex int开始的下标
path []int保持路径上的一组解
result *[][]int保存总的结果集

*/
func backtrack(candidates []int, target int, sum int, startIndex int, path []int, result *[][]int) {
	//结束条件  sum == target
	if sum > target { //剪枝函数
		return
	}

	if sum == target {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)
		return
	}

	//单层for循环
	for i := startIndex; i < len(candidates) && sum+candidates[i] <= target; i++ {
		sum += candidates[i] //加上数组的元素
		path = append(path, candidates[i])
		//递归遍历
		fmt.Println("回溯前：", path)
		backtrack(candidates, target, sum, i, path, result) //可以选取重复元素 i
		// 回溯
		sum -= candidates[i]
		path = path[:len(path)-1]
		fmt.Println("回溯后", path)
	}
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates) //按顺序排序
	path := make([]int, 0)
	result := make([][]int, 0)
	backtrack(candidates, target, 0, 0, path, &result)
	return result
}

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	result := combinationSum(candidates, target)
	fmt.Println(result)
}
