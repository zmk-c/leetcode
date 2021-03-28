package main

import (
	"fmt"
	"sort"
)

/*与leet78_子集问题的区别:
	leet78_子集中的nums数组不包含重复数字，本题中所给的nums数字可能包含重复数字，而且要求最后的结果集中不包含重复子集
涉及到了集合去重，这个去重与leet40_组合总和II类似
*/

/*参数确定：
nums []int 题目所给数组
path []int 一组解
startIndex int 开始遍历的位置
used []bool 用于判断同一树层是否使用过，排除重复的子集
result *[][]int 最终解集
*/
func backtrack(nums, path []int, startIndex int, used []bool, result *[][]int) {
	tmp := make([]int, len(path))
	copy(tmp, path)
	*result = append(*result, tmp)

	//单层for循环判断
	for i := startIndex; i < len(nums); i++ {
		//nums[i-1] == nums[i]&&used[i-1]==false表示同一树层nums[i-1]已经使用过了
		//nums[i-1] == nums[i]&&used[i-1]==true表示同一树枝nums[i-1]已经使用过了
		if i > 0 && nums[i-1] == nums[i] && used[i-1] == false {
			continue
		}
		path = append(path, nums[i])
		used[i] = true
		backtrack(nums, path, i+1, used, result)
		used[i] = false
		path = path[:len(path)-1]
	}

}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums) //需要先排序
	path := make([]int, 0)
	used := make([]bool, len(nums))
	result := make([][]int, 0)
	backtrack(nums, path, 0, used, &result)
	return result
}

func main() {
	nums := []int{1, 2, 2}
	result := subsetsWithDup(nums)
	fmt.Println(result)
}
