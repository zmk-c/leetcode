package main

import (
	"fmt"
	"sort"
)

/*
给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。

示例 1：
输入：nums = [1,1,2]
输出：
[[1,1,2],
 [1,2,1],
 [2,1,1]]

示例 2：
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]


提示：
1 <= nums.length <= 8
-10 <= nums[i] <= 10

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*参数确定
nums []int 题目所给数组
path []int 一组解
used []bool 进行去重操作与判断该元素是否使用过
result *[][]int 最终解集
*/
func backtrack(nums, path []int, used []bool, result *[][]int) {
	//终止条件 path中的元素个数=nums中的个数
	if len(path) == len(nums) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)
		return
	}

	//单层for循环判断
	for i := 0; i < len(nums); i++ { //全排列 从开头遍历used数组中的元素是否被访问过
		if i > 0 && nums[i-1] == nums[i] && used[i-1] == false {
			continue
		}
		if used[i] == false { //判断当前元素是否使用过
			used[i] = true
			path = append(path, nums[i])
			backtrack(nums, path, used, result)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
}

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	path := make([]int, 0)
	used := make([]bool, len(nums))
	result := make([][]int, 0)
	backtrack(nums, path, used, &result)
	return result
}

func main() {
	nums := []int{1, 1, 2}
	result := permuteUnique(nums)
	fmt.Println(result)
}
