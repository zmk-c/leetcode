package main

import "fmt"

func SelectSort(array []int) {
	n := len(array)
	for i := 0; i < n; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if array[minIndex] > array[j] {
				minIndex = j
			}
		}
		array[i], array[minIndex] = array[minIndex], array[i]
	}
}

func main() {
	array := []int{1, 4, 2, 1, 3}
	SelectSort(array)
	fmt.Println("冒泡排序后的数组为：", array)
}
