package main

import "fmt"

/*思路：按照题目的要求进行模拟，我们在遍历字符串的时候，令i+=(2*k)，每次增加的区间就是2k个字符串，
然后再判断是否有需要反转的区间。*/
func reverseStr(s string, k int) string {

	strs := []byte(s)
	for i := 0; i < len(strs); i += (2 * k) {
		//如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样
		l := i //不能继续用变量i  否则后面的for循环会改变局部变量的值
		r := i + k - 1
		if r >= len(strs) {
			//如果剩余字符少于 k 个，则将剩余字符全部反转。
			r = len(strs) - 1
		}

		for l < r && l < len(strs) {
			strs[l], strs[r] = strs[r], strs[l]
			l++
			r--
		}
	}
	return string(strs)
}

func main() {
	s := "abcdefg"
	k := 2
	revStr := reverseStr(s, k)
	fmt.Println(revStr)
}
