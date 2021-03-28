package main

import (
	"fmt"
	"sort"
)

/*
参数说明：
candidates []int 题目提供的数组
target int 题目提供的目标和
sum int path路径上的累加和 需要与target进行比较
startIndex int 开始搜索的下标
used []bool 用于去重，used[i-1]==true说明同一树枝的candidates[i-1]使用过，used[i-1]==false说明同一树层的candidates[i-1]使用过
path []int 一组解
result *[][]int 最终的结果集
*/
func backtrack(candidates []int, target int, sum int, startIndex int, used []bool, path []int, result *[][]int) {
	//判断退出条件
	if sum == target {
		//加入结果集中
		tmp := make([]int, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)
		return
	}

	//单个for循环的判断
	for i := startIndex; i < len(candidates) && sum+candidates[i] <= target; i++ {
		//used[i-1]==true说明同一树枝的candidates[i-1]使用过
		// used[i-1]==false说明同一树层的candidates[i-1]使用过
		//要对同一树层的使用过的元素进行跳过
		if i > 0 && candidates[i] == candidates[i-1] && used[i-1] == false {
			continue
		}
		sum += candidates[i]
		path = append(path, candidates[i])
		used[i] = true
		fmt.Println("回溯前：", path)
		backtrack(candidates, target, sum, i+1, used, path, result) //每个数字在集合中只能使用一次
		used[i] = false
		sum -= candidates[i]
		path = path[:len(path)-1]
		fmt.Println("回溯后：", path)
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	used := make([]bool, len(candidates))
	path := make([]int, 0)
	result := make([][]int, 0)
	backtrack(candidates, target, 0, 0, used, path, &result)
	return result
}

func main() {
	candidates := []int{2, 5, 2, 1, 2}
	target := 5
	result := combinationSum2(candidates, target)
	fmt.Println(result)
}
