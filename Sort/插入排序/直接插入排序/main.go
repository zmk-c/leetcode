package main

import "fmt"

/*
直接插入排序：简单来说，就是在每一步将一个待排序的数据，从后向前找到合适的位置插入，直到全部数据插入完，排序结束。
即将array[i]插入前面array[1]~array[i-1]已经有序的队列中
（从第二个数开始，第一个数插入时没有其他数据，所以直接不用管）
*/
func DirectInsertSort(array []int) {
	n := len(array)
	for i := 1; i < n; i++ { //第一个元素是已经有序的
		key := array[i] //x为临时变量 作用为备份当前元素
		j := i - 1      //从前一个元素开始比较
		for {
			if j > 0 && array[j] > key { //表示while循环
				array[j+1] = array[j]
				j = j - 1
			} else {
				break
			}
		}
		array[j+1] = key
	}
}

func main() {
	array := []int{1, 4, 2, 1, 3}
	DirectInsertSort(array)
	fmt.Println("直接插入排序后的数组为：", array)
}
