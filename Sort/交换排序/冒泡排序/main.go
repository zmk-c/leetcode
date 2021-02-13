package main

import "fmt"

func BubbleSort(array []int) {
	n := len(array)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if array[j+1] <= array[j] {
				tmp := array[j]
				array[j] = array[j+1]
				array[j+1] = tmp
			}
		}
	}
}

func main() {
	array := []int{1, 4, 2, 1, 3}
	BubbleSort(array)
	fmt.Println("冒泡排序后的数组为：", array)
}
