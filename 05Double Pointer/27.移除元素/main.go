package main

/*1.这个题目暴力的解法就是两层for循环，一个for循环遍历数组元素 ，第二个for循环更新数组。
时间复杂度：O(n2)
空间复杂度：O(1)
*/
// func removeElement(nums []int, val int) int {
// 	length := len(nums)
// 	for i := 0; i < length; i++ {
// 		if nums[i] == val { //找到了目标值，将后续的数组内容往前进行覆盖
// 			for j := i + 1; j < length; j++ {
// 				nums[j-1] = nums[j]
// 			}
// 			i--
// 			length-- //数组长度-1
// 		}
// 	}
// 	return length
// }

/*
2.双指针法（快慢指针法）：「通过一个快指针和慢指针在一个for循环下完成两个for循环的工作。」
双指针法在数组和链表的操作中是非常常见的，很多考察数组和链表操作的面试题，都使用双指针法。
*/
func removeElement(nums []int, val int) int {
	slowIndex := 0
	for fastIndex := 0; fastIndex < len(nums); fastIndex++ {
		if val != nums[fastIndex] {
			nums[slowIndex] = nums[fastIndex]
			slowIndex++
		}
	}
	return slowIndex
}
