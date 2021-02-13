package main

import "fmt"

/*
给定一个整型数组, 你的任务是找到所有该数组的递增子序列，递增子序列的长度至少是2。

示例:

输入: [4, 6, 7, 7]
输出: [[4, 6], [4, 7], [4, 6, 7], [4, 6, 7, 7], [6, 7], [6, 7, 7], [7,7], [4,7,7]]
说明:

给定数组的长度不会超过15。
数组中的整数范围是 [-100,100]。
给定数组中可能包含重复数字，相等的数字应该被视为递增的一种情况。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/increasing-subsequences
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*参数确定：
1.nums []int 题目所给的数组
2.startIndex int 遍历的起始位置，元素不能重复使用
3.path []int 一组解
4.result *[][]int 最终解集
*/
func backtrack(nums []int, startIndex int, path []int, result *[][]int) {
	//终止条件
	if len(path) >= 2 { //需要遍历整棵树并且递增子序列的长度至少为2
		tmp := make([]int, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)
	}

	//单层for循环操作
	//利用map进行去重操作
	used := make(map[int]bool)
	for i := startIndex; i < len(nums); i++ {
		//如果选取元素num[i]小于path中最后一位数 或者是 同一层使用了重复的元素，则跳过
		if len(path) != 0 && nums[i] < path[len(path)-1] || used[nums[i]] == true {
			continue
		}
		used[nums[i]] = true
		path = append(path, nums[i])
		backtrack(nums, i+1, path, result)
		path = path[:len(path)-1]
	}

}

func findSubsequences(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	result := make([][]int, 0)
	backtrack(nums, 0, []int{}, &result)
	return result
}

func main() {
	nums := []int{4, 6, 7, 7}
	result := findSubsequences(nums)
	fmt.Println(result)
}
