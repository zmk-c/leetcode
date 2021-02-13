package main

import (
	"fmt"
	"sort"
)

/*
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。
解集不能包含重复的组合。
示例 1：

输入：candidates = [2,3,6,7], target = 7,
所求解集为：
[
  [7],
  [2,2,3]
]
示例 2：

输入：candidates = [2,3,5], target = 8,
所求解集为：
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]


提示：

1 <= candidates.length <= 30
1 <= candidates[i] <= 200
candidate 中的每个元素都是独一无二的。
1 <= target <= 500

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combination-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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
