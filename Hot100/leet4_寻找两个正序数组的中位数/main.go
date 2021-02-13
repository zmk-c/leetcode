package main

import "fmt"

/*
给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的中位数。

进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？

示例 1：

输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
示例 2：

输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
示例 3：

输入：nums1 = [0,0], nums2 = [0,0]
输出：0.00000
示例 4：

输入：nums1 = [], nums2 = [1]
输出：1.00000
示例 5：

输入：nums1 = [2], nums2 = []
输出：2.00000
 

提示：

nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
1 <= m + n <= 2000
-106 <= nums1[i], nums2[i] <= 106

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*最暴力的方法进行两个数组的合并排序，然后取中间两个数进行计算中位数。时间复杂度O(m+n) 由于开辟的一个新的数组，空间复杂度O(m+n) */
//但是达不到题目所给的O(log(m+n))的时间度

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	var result float64

	//如果有一个数组为空  则从另一个数组中选取中位数    这里不考虑两个数组都为空的情况
	if l1==0 && l2>0 {
		result = getResult(l2,nums2)
	}
	if l2==0 && l1>0 {
		result = getResult(l1,nums1)
	}

	//两个数组都不为空的情况
	//将nums1的前m项与nums2的前n项进行合并
	tmp := make([]int,l1+l2)
	n1,n2,t := 0,0,0
	for n1<l1 && n2<l2{
		if nums1[n1] <= nums2[n2]{  //将nums1中的数组与nums2数组进行比较  将较小的一方加入到tmp中
			tmp[t] = nums1[n1]
			t++
			n1++
		}else {
			tmp[t] = nums2[n2]
			t++
			n2++
		}
	}

	//如果nums1数组中还有剩余  则将其加入tmp中
	for n1 < l1{
		tmp[t] = nums1[n1]
		t++
		n1++
	}

	//对于nums2同理
	for n2 < l2{
		tmp[t] = nums2[n2]
		t++
		n2++
	}

	result = getResult(len(tmp),tmp)

	return result
}

func getResult(length int,arr []int) float64{
	var result float64
	if length % 2 == 0 {
		result = float64(arr[length/2] + arr[length/2 - 1]) / float64(2) //长度位偶数位计算中位数
	}else {
		result = float64(arr[length/2])	//长度位奇数计算中位数
	}
	return result
}

func main() {
	nums1 := []int{2}
	//nums1 := []int{0,0}
	//nums1 := []int{}
	nums2 := []int{1,3}
	result := findMedianSortedArrays(nums1, nums2)
	fmt.Println(result)
}
