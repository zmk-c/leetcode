package main

import (
	"fmt"
	"math"
)

/*
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
返回滑动窗口中的最大值。

示例 1：
输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7

示例 2：
输入：nums = [1], k = 1
输出：[1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sliding-window-maximum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//暴力超时
func maxSlidingWindowBL(nums []int, k int) []int {
	queue := make([]int, k)
	result := make([]int, 0)
	for i := 0; i < k; i++ {
		queue[i] = nums[i]
	}

	for i := k; i <= len(nums); i++ {
		max := getMax(queue)
		result = append(result, max)
		queue = queue[1:]
		if i == len(nums) {
			break
		}
		queue = append(queue, nums[i])
	}
	return result
}

func getMax(arr []int) int {
	var max int = math.MinInt64 //一开始初始值为0有问题
	for _, v := range arr {
		if v >= max {
			max = v
		}
	}
	return max
}

//---------分割线---------
/*双端单调递减队列*/
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	//用切片模拟一个双端队列
	queue := []int{}
	result := []int{} //保存结果
	for i := range nums {
		/*push()函数操作*/
		//遍历nums[i]和队列尾部元素比较，如果大于队列尾部元素 则将队列尾部元素pop出队
		for len(queue) > 0 && nums[i] > queue[len(queue)-1] {
			//将比当前nums[i]小的元素都pop掉
			queue = queue[:len(queue)-1]
		}
		//然后将当前元素放入queue中  可能是窗口内元素的最大值
		queue = append(queue, nums[i])

		/*pop()函数操作*/
		//每次弹出的时候，比较当前要弹出的数值是否等于队列头部出口元素的数值，如果相等则弹出。
		if i >= k && nums[i-k] == queue[0] {
			queue = queue[1:]
		}

		if i >= k-1 {
			//放入结果数组
			result = append(result, queue[0])
		}
	}
	return result
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	result := maxSlidingWindow(nums, 3)
	fmt.Println(result)
}
