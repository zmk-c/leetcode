package main

import (
	"math"
)

//暴力解法：利用两层for循环 然后不断寻找符合条件的子序列，判断符合子序列长度的最小值   O(n^2)
func minSubArrayLenBL(s int, nums []int) int {
	var sum int                //记录遍历过程子序列的值  需要与s比较
	subLength := math.MaxInt32 //子序列的长度 如果返回0表示不符合条件
	for i := 0; i < len(nums); i++ {
		sum = 0 //每次寻找时都要使sum置0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum >= s { //结束条件  子序列的和sum大于目标值s
				subLength = min(subLength, j-i+1) ////取最小的满足条件的子序列长度
				break
			}
		}
	}
	if subLength == math.MaxInt32 { //如果subLength没有被赋值  说明没有适合的子序列 返回0
		return 0
	}
	return subLength
}

/*使用滑动窗口
所谓滑动窗口，「就是不断的调节子序列的起始位置和终止位置，从而得出我们要想的结果」

在本题中实现滑动窗口，主要确定如下三点：
窗口内是什么？  => 满足其和 ≥ s 的长度最小的 连续 子数组
如何移动窗口的起始位置？  => 如果当前窗口的值大于s了，窗口就要向前移动了
如何移动窗口的结束位置？  => 窗口的结束位置就是遍历数组的指针，窗口的起始位置设置为数组的起始位置就可以了。

「滑动窗口的精妙之处在于根据当前子序列和大小的情况，不断调节子序列的起始位置。从而将O(n^2)的暴力解法降为O(n)。」
*/
func minSubArrayLen(s int, nums []int) int {
	var (
		sum, i, j = 0, 0, 0       //sum为滑动窗口数值之和 需要与s比较, i为滑动窗口的开始位,j位滑动窗口的结束位
		subLength = math.MaxInt32 //滑动窗口长度(子序列的长度) 如果返回0表示不符合条件
	)
	for j < len(nums) {
		sum += nums[j]
		for sum >= s {
			subLength = min(subLength, j-i+1)
			sum -= nums[i]
			i++
		}
		j++
	}
	if subLength == math.MaxInt32 { //如果subLength没有被赋值  说明没有适合的子序列 返回0
		return 0
	}
	return subLength
}

//辅助函数
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
