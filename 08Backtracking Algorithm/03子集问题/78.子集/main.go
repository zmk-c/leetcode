package main

import "fmt"

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
