package main

import "fmt"

/*
编写一个算法来判断一个数 n 是不是快乐数。
「快乐数」定义为：对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和，然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。如果 可以变为  1，那么这个数就是快乐数。
如果 n 是快乐数就返回 True ；不是，则返回 False 。


示例：
输入：19
输出：true
解释：
12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/happy-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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
		} else {
			mp[sum] = true
		}
		n = sum
	}
}

func main() {
	n := 19
	isOne := isHappy(n)
	fmt.Println(isOne)
}
