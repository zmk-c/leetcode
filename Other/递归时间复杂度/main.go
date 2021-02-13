package main

import "fmt"

//面试题：求x的n次方
//并计算其时间复杂度

//1.O(N)时间复杂度
func cal01(x int, n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result = result * x
	}
	return result
}

//1.O(N)时间复杂度
func cal02(x int, n int) int {
	if n == 0 {
		return 1
	}
	return cal01(x, n-1) * x
}

//1.O(N)时间复杂度
func cal03(x int, n int) int {
	if n == 0 {
		return 1
	}

	if n%2 == 1 {
		return cal03(x, n/2) * cal03(x, n/2) * x
	} else {
		return cal03(x, n/2) * cal02(x, n/2)
	}
}

//2.O(logn)时间复杂度
func cal04(x int, n int) int {
	if n == 0 {
		return 1
	}

	t := cal04(x, n/2)
	if n%2 == 1 {
		return t * t * x
	} else {
		return t * t
	}
}

func main() {
	result1 := cal01(2, 6)
	fmt.Println(result1) //64

	result2 := cal02(2, 7)
	fmt.Println(result2) //128

	result3 := cal03(2, 8)
	fmt.Println(result3) //256

	result4 := cal04(2, 9)
	fmt.Println(result4) //512
}
