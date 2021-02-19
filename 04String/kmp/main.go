package main

import "fmt"

//当前字符串冲突时回退到前一位next数组对应的值
func getNext(next []int, s string) {
	//初始化
	j := 0
	next[0] = 0
	for i := 1; i < len(s); i++ {
		//前后缀不相同的情况
		for j > 0 && s[i] != s[j] {
			j = next[j-1] //回退
		}
		//前后缀相同的情况
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
}

func main() {
	s := "aabaaf"
	next := make([]int, len(s))
	getNext(next, s)
	for i := 0; i < len(s); i++ {
		fmt.Println(next[i])
	}
}
