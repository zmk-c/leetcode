package main

import (
	"fmt"
)

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

示例:

给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*暴力枚举*/
/*
时间复杂度：O(N^2)，其中 NN 是数组中的元素数量。最坏情况下数组中任意两个数都要被匹配一次。
空间复杂度：O(1)
我的考虑：
	1.假设nums[0]+x=target
	2.x=target-num[0]
	3.遍历数组是否包含x，如果包含则返回其下标
*/
/*func TwoSum(nums []int, target int) []int {
	var x int
	result := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		x = target - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if x == nums[j] {
				result = append(result, i, j)
				break
			}
		}
	}
	return result
}*/

/*哈希表*/
//哈希表中的key为nums中的值，val为值的索引下标
func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for i, num := range nums {
		another := target - num //与num和为target的另一个数
		if anotherIndex, ok := mp[another]; ok {
			return []int{anotherIndex, i}
		} else {
			mp[num] = i
		}
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result)
}
