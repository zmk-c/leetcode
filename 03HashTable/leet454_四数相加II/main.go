package main

import "fmt"

/*
给定四个包含整数的数组列表 A , B , C , D ,计算有多少个元组 (i, j, k, l) ，使得 A[i] + B[j] + C[k] + D[l] = 0。
为了使问题简单化，所有的 A, B, C, D 具有相同的长度 N，且 0 ≤ N ≤ 500 。所有整数的范围在 -228 到 228 - 1 之间，最终结果不会超过 231 - 1 。

例如:
输入:
A = [ 1, 2]
B = [-2,-1]
C = [-1, 2]
D = [ 0, 2]

输出:
2

解释:
两个元组如下:
1. (0, 0, 0, 1) -> A[0] + B[0] + C[0] + D[1] = 1 + (-2) + (-1) + 2 = 0
2. (1, 1, 0, 0) -> A[1] + B[1] + C[0] + D[0] = 2 + (-1) + (-1) + 0 = 0

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/4sum-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
解题步骤：
1.首先定义 一个map，key放a和b两数之和，value 放a和b两数之和出现的次数。
2.遍历大A和大B数组，统计两个数组元素之和，和出现的次数，放到map中。
3.定义int变量count，用来统计a+b+c+d = 0 出现的次数。
4.在遍历大C和大D数组，找到如果 0-(c+d) 简化为-c-d 在map中出现过的话，就用count把map中key对应的value也就是出现次数统计出来。
5.最后返回统计值 count 就可以了
*/
func fourSumCount(A []int, B []int, C []int, D []int) int {
	mp := make(map[int]int, len(A)*len(B))
	count := 0
	for _, a := range A {
		for _, b := range B {
			mp[a+b]++
		}
	}

	for _, c := range C {
		for _, d := range D {
			if n, ok := mp[-c-d]; ok {
				count += n //n为mp[-c-d]也就是mp[a+b]和的次数
			}
		}
	}

	return count
}

func main() {
	A := []int{1, 2}
	B := []int{-2, -1}
	C := []int{-1, 2}
	D := []int{0, 2}
	count := fourSumCount(A, B, C, D)

	fmt.Println(count)
}
