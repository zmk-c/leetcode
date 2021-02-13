package main

import "fmt"

/*
给你一个整数数组 nums ，返回该数组所有可能的子集（幂集）。解集不能包含重复的子集。

示例 1：
输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

示例 2：
输入：nums = [0]
输出：[[],[0]]

提示：
1 <= nums.length <= 10
-10 <= nums[i] <= 10

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subsets
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*参数确定：
nums []int 题目所提供的数组
startIndex int 开始遍历的下标
path []int 一组解
result *[][]int 最终结果集
*/
func backtrack(nums, path []int, startIndex int, result *[][]int) {
	//终止条件  startIndex >= len(nums) 可以不添加，遍历完自带终止了
	tmp := make([]int, len(path))
	copy(tmp, path)
	*result = append(*result, tmp)

	//单层for循环
	for i := startIndex; i < len(nums); i++ { //子集合具有无序性，取过的元素不会重复，for循环以startIndex开始
		path = append(path, nums[i])
		//fmt.Println("回溯前:", path)
		backtrack(nums, path, i+1, result)
		path = path[:len(path)-1]
		//fmt.Println("回溯后:", path)
	}
}

func subsets(nums []int) [][]int {
	path := make([]int, 0)
	result := make([][]int, 0)
	backtrack(nums, path, 0, &result)
	return result
}

func main() {
	nums := []int{1, 2, 3}
	result := subsets(nums)
	fmt.Println(result)
}
