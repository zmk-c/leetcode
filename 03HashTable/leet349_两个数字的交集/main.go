package main

import "fmt"

/*
给定两个数组，编写一个函数来计算它们的交集。

示例 1：
输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2]

示例 2：
输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[9,4]

说明：
输出结果中的每个元素一定是唯一的。
我们可以不考虑输出结果的顺序。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/intersection-of-two-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//可以利用map实现set的集合 map集合的key作为唯一值  并实现实现去重
/*func intersection(nums1 []int, nums2 []int) []int {
	mp := make(map[int]bool)
	result := make([]int, 0)

	for _, n1 := range nums1 { //将nums1中的数作为mp中的键，实现去重效果
		mp[n1] = true
	}

	for _, n2 := range nums2 {
		if mp[n2] {
			mp[n2] = false
			result = append(result, n2)
		}
	}

	return result
}*/

//map的value值是布尔型，这会导致set多占用内存空间，解决这个问题，则可以将其替换为空结构。在Go中，空结构通常不使用任何内存。
//unsafe.Sizeof(struct{}{}) // 结果为 0
func intersection(nums1 []int, nums2 []int) []int {
	mp := make(map[int]struct{})
	result := make([]int, 0)

	for _, n1 := range nums1 { //将nums1中的数作为mp中的键，实现去重效果
		mp[n1] = struct{}{}
		fmt.Println(mp)
	}

	for _, n2 := range nums2 {
		if _, ok := mp[n2]; ok {
			result = append(result, n2)
			delete(mp, n2)
		}
	}

	return result
}

func main() {
	num1 := []int{1, 2, 2, 1}
	num2 := []int{2, 2}
	result := intersection(num1, num2)
	fmt.Println(result)
}
