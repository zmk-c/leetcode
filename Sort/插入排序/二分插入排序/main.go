package main

import "fmt"

/*思路和直接插入排序类似，不过插入的位置用二分查找进行了优化，由原来的O(n)时间变为了O（logn) 所以总时间位O(nlogn)*/
func BinaryInsertSort(array []int) {
	n := len(array)
	low, high, mid, key := 0, 0, 0, 0
	for i := 0; i < n; i++ {
		key = array[i]
		low = 0
		high = i - 1
		for { //二分查找到位置
			if low <= high {
				mid = (low + high) / 2
				if key < array[mid] {
					high = mid - 1
				} else {
					low = mid + 1
				}
			} else {
				break
			}
		}
		// 找到位置后需要将原先位置上及后面的的数据向后平移一位，腾出位置插入
		for j := i - 1; j >= low; j-- {
			array[j+1] = array[j]
		}
		//将元素插入
		if low != i { //low == i 表示该元素就在所应当在的位置上，无需排序
			array[low] = key
		}
	}

}

func main() {
	array := []int{1, 4, 2, 1, 3}
	BinaryInsertSort(array)
	fmt.Println("二分插入排序后的数组为：", array)
}
