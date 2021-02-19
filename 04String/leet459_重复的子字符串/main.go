package main

import "fmt"

/*
题目459.重复的子字符串
给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。给定的字符串只含有小写英文字母，并且长度不超过10000。

示例 1:
输入: "abab"
输出: True
解释: 可由子字符串 "ab" 重复两次构成。

示例 2:
输入: "aba"
输出: False

示例 3:
输入: "abcabcabcabc"
输出: True
解释: 可由子字符串 "abc" 重复四次构成。(或者子字符串 "abcabc" 重复两次构成。)
*/
//构造前缀表
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

func repeatedSubstringPattern(s string) bool {
	l := len(s)
	if l == 0 {
		return false
	}
	next := make([]int, l)
	getNext(next, s)

	//如果字符串的长度-最长相等前后缀长度  能被字符串长度除尽  则说明其子串能够构成该字符串
	//且最长相等前后缀长度不为0
	if next[l-1] != 0 && l%(l-next[l-1]) == 0 {
		return true
	}
	return false
}

func main() {
	s := "abcabcabc"
	flag := repeatedSubstringPattern(s)
	fmt.Println(flag)
}
