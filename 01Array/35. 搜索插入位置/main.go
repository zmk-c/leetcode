package main

/*
无非四种情况：
目标值在数组所有元素之前
目标值等于数组中某一个元素
目标值插入数组中的位置
目标值在数组所有元素之后

因为是有序数组，可以利用二分查找
二分查找涉及的很多的边界条件，逻辑比较简单，就是写不好。
例如到底是 while(left < right) 还是 while(left <= right)，到底是right = middle呢，还是要right = middle - 1呢？

对于[left,right]左闭右闭的情况  选择 while(left <= right) , right = middle -1
对于[left,right)左闭右开的情况，选择 while(left < right) , right = middle
*/
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right { //[left,right]  left==right有效
		//middle := (left + right) / 2
		middle := left + ((right - left) >> 1) //防止整数溢出 并利用移位提高速度
		if nums[middle] > target {
			right = middle - 1
		} else if nums[middle] < target {
			left = middle + 1
		} else { //此时target == num[middle]  返回索引
			return middle
		}
	}
	return right + 1
}
