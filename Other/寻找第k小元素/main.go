package main

import (
	"fmt"
	"sort"
	"time"
)

/*最先想到的就使将这n个元素先从小到大进行排序，然后在取出第k个元素，即为目标元素  时间复杂度为O(nlogn)
当数量很大时，效果变差
 */
/*func findTheKstNum(arr []int,k int) int{
	sort.Ints(arr) //golang中自带排序的函数
	fmt.Println("sorted arr:",arr)
	return arr[k]
}*/

/*
采用分治的方法进行改进:
	(1) 当规模小于阈值（n<44）时，直接用排序算法返回结果。
	(2) 当n大于阈值时，把n个元素划分为5个元素一组的n/5组，排除剩余元素（不会有影响，这里只是为了求中项mid），
	分别排序，然后挑出每一组元素的中间值，再在所有的中间值中，递归调用本算法，挑出中间值mid。
	(3) 把元素划分为A1、A2、A3三组，分别包含小于、等于、大于mid的元素。
	(4)分三种情况：
		若A1的元素数量大于等于K，即第K个元素在第一组内：在A1中递归查找第k小元素。
		若A1、A2元素个数之和大于等于K，即中项mid为第K个元素：返回mid
		否则，第K个元素在第三组：在A3中递归寻找第（k-|A1、A2元素数量之和|）小元素。
*/

func findNum(arr []int,k int) int{
	return findTheKstNum(arr,0,len(arr)-1,k)
}

func findTheKstNum(arr []int,low int,high int,k int) int{
	n := high - low + 1
	//当规模小于阈值（n<44）时，直接用排序算法返回结果。
	if n < 44{
		sort.Ints(arr)
		fmt.Println("sorted:",arr)
		return arr[k]
	}
	//当n大于阈值时，把n个元素划分为5个元素一组的n/5组，再排序
	q := n / 5
	M := make([]int,q)
	for i:=0;i<q;i++{
		sort.Ints(arr[low+i*5:low+i*5+5])
		M[i] = arr[low+i*5+2] //挑出每一组元素的中间值加入中间值数组
	}
	fmt.Println("M:",M)
	//再在所有的中间值中，递归调用本算法，挑出中间值mid
	mid := findTheKstNum(M,0,q,q/2)
	// 把元素划分为A1、A2、A3三组，分别包含小于、等于、大于mid的元素
	A1 := make([]int,0)
	A2 := make([]int,0)
	A3 := make([]int,0)
	for i:=low;i<high;i++{
		if arr[i] < mid{
			A1 = append(A1,arr[i])
		} else if arr[i] == mid{
			A2 = append(A2,arr[i])
		} else {
			A3 = append(A3,arr[i])
		}
	}
	fmt.Println("A1:",A1)
	fmt.Println("A2:",A2)
	fmt.Println("A3:",A3)
	if len(A1) > k { //如果小于A1的数量大于k,则第k小的元素在A1数组中
		return findTheKstNum(A1,0,len(A1)-1,k)
	}else if len(A1)+len(A2) >= k {	//如果A1和A2的元素个数之和大于k,则mid即为第k小的值
		return mid
	}else {	//否则，第k小的数在A3集合中，在A3集合中递归遍历到k-len(A1)-len(A2)的数
		return findTheKstNum(A3,0,len(A3),k - len(A1) - len(A2))
	}
}

func main(){
	//arr := []int{4,5,12,1,12,4,1,0,9}
	//arr := []int{9,12,34,5,13,4,5,6,2,1,0,9,12,34,5,13,4,5,6,2,1,0,9,12,34,5,13,4,5,6,2,1,0,9,12,34,5,13,4,5,6,2,1,0,9,12,34,5,13,4,5,6,2,1,0,9,12,34,5,13,4,5,6,2,1,0}
	arr := []int{0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45}
	k := 9
	start := time.Now().UnixNano()
	//num := findTheKstNum(arr,k)
	num := findNum(arr, k)
	end := time.Now().UnixNano()
	fmt.Printf("第%d个元素为%d \n",k,num)
	fmt.Printf("所花费的时间为:%d ns",end - start)
}
