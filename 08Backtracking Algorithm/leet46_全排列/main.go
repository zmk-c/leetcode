package main

import "fmt"

/*
注意：
也就是说如果切片本身的容量不足以容纳追加后的所有元素，那么会申请一个新的底层数组（切片是一种基于底层数组的数据实现）。
这也就是为什么我们在使用append()函数的时候，总是需要接收其返回值。因为我们传进去的切片变量的地址是可能发生变化的。

解决方法 通过每一步显式返回res来解决append()导致res地址变化的问题：\
		还可以利用指针

也可以采用全局变量的方式设置result数组
*/

/*
回溯核心
nums: 原始数据列表
path: 路径上的数字
used: 是否访问过

result: 返回的结果数组
*/
func backtrack(nums, path []int, used []bool, result *[][]int) {
	// 1.结束条件：走完了，也就是路径上的数字总数等于原始列表总数
	if len(nums) == len(path) {
		tmp := make([]int, len(nums))
		// 切片底层公用数据，所以要copy
		copy(tmp, path)
		// 把本次排列结果追加到最终结果上
		*result = append(*result, tmp)
		return
	}

	// 开始遍历原始数组的每个数字
	for i := 0; i < len(nums); i++ {
		// 检查是否访问过
		if !used[i] {
			// 没有访问过就选择它，然后标记成已访问过的
			used[i] = true
			//2.做选择：将这个数字加入到当前排列
			path = append(path, nums[i])
			backtrack(nums, path, used, result)
			//3.撤销刚才的选择，也就是恢复操作
			path = path[:len(path)-1]
			// 标记成未使用
			used[i] = false
		}
	}
}

func permute(nums []int) [][]int {
	path := make([]int, 0)
	used := make([]bool, len(nums))
	result := make([][]int, 0)
	backtrack(nums, path, used, &result)
	return result
}

func main() {
	nums := []int{1, 2, 3}
	result := permute(nums)
	fmt.Println("全排列结果为：", result)
}
