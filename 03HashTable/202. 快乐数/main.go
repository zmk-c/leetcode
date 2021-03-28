package main

import "fmt"

//使用哈希算法来判断这个sum是否重复出现，如果重复出现，返回false。否则一直找到sum=1为止
func getSum(n int) int {
	sum := 0
	for n != 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}

func isHappy(n int) bool {
	mp := make(map[int]bool)
	for {
		sum := getSum(n)
		if sum == 1 { //如果sum==1则表明找到快乐数  退出
			return true
		}

		//如果这个sum曾经出现过  则表明陷入死循环了  return false
		if mp[sum] == true {
			return false
		}
		mp[sum] = true

		n = sum
	}
}

func main() {
	n := 19
	isOne := isHappy(n)
	fmt.Println(isOne)
}
