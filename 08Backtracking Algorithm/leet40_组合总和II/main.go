package main

import (
	"fmt"
	"sort"
)

/*
给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
candidates 中的每个数字在每个组合中只能使用一次。

说明：

所有数字（包括目标数）都是正整数。
解集不能包含重复的组合。
示例 1:

输入: candidates = [10,1,2,7,6,1,5], target = 8,
所求解集为:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
示例 2:

输入: candidates = [2,5,2,1,2], target = 5,
所求解集为:
[
  [1,2,2],
  [5]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combination-sum-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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
