package main

import "fmt"

/*
活动安排问题就是要在所给的活动集合中选出最大的相容活动子集合，这是可以用贪心算法有效解决的一个例子
	每个活动都包含开始时间和结束时间，即每个活动的活动时间区间为[si,fi)，贪心算法选择总是最早完成时间的活动加入安排集合
使剩余的可安排时间最大化。先将活动按照最早的结束时间进行升序排序，若后一个活动的开始时间>=前一个活动的结束时间，则加入安排集合。
依次类推
*/

//参数：每个活动的开始时间与结束时间
//返回值：被选中的活动总数和每个活动是否被选择
func GreedySelector(start []int, finish []int) (int, []bool) {
	//1.先将活动按照最早的结束时间进行升序排序 并让活动开始时间与之对应
	n := len(start)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if finish[i] >= finish[j] {
				tmp := finish[i]
				finish[i] = finish[j]
				finish[j] = tmp

				h := start[i]
				start[i] = start[j]
				start[j] = h
			}
		}
	}

	//2.选择活动
	selector := make([]bool, n)
	selector[0] = true  //最开始的活动肯定被入选
	var j, count = 0, 1 //count表示计算被选的活动的总数
	for i := 1; i < n; i++ {
		if start[i] > finish[j] { //若后一个活动的开始时间>=前一个活动的结束时间，则加入安排集合
			selector[i] = true
			j = i
			count++
		} else {
			selector[i] = false
		}
	}

	return count, selector
}

func main() {
	start := []int{12, 2, 8, 8, 6, 5, 3, 5, 0, 3, 1}
	finish := []int{14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4}
	count, selector := GreedySelector(start, finish)
	fmt.Printf("活动选择情况如下：%v \n 共选了%v个活动", selector, count)
}
