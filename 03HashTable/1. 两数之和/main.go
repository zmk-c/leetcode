package main

/*暴力枚举*/
/*
时间复杂度：O(N^2)，其中 NN 是数组中的元素数量。最坏情况下数组中任意两个数都要被匹配一次。
空间复杂度：O(1)
我的考虑：
	1.假设nums[0]+x=target
	2.x=target-num[0]
	3.遍历数组是否包含x，如果包含则返回其下标
*/
// func TwoSum(nums []int, target int) []int {
// 	var x int
// 	result := make([]int, 0)
// 	for i := 0; i < len(nums); i++ {
// 		x = target - nums[i]
// 		for j := i + 1; j < len(nums); j++ {
// 			if x == nums[j] {
// 				result = append(result, i, j)
// 				break
// 			}
// 		}
// 	}
// 	return result
// }

/*哈希表*/
//哈希表中的key为nums中的值，val为值的索引下标
func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for i, num := range nums {
		another := target - num //与num和为target的另一个数
		if anotherIndex, ok := mp[another]; ok {
			return []int{anotherIndex, i}
		}
		mp[num] = i
	}
	return []int{}
}
