package main

import "fmt"

//求解next[]数组
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

//进行匹配：为了方便，将haystack字符串称为文本串，将needle字符串称为模式串
func strStr(haystack string, needle string) int {
	//如果模式串为空串 则返回0
	if len(needle) == 0 {
		return 0
	}

	//求解模式串的next数组
	next := make([]int, len(needle))
	getNext(next, needle)

	j := 0 // next[]数组中记录的起始位置为0

	//开始匹配
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] { //字符不匹配的情况
			j = next[j-1] //这里j要找前一位回退的位置
		}
		//字符匹配情况，i,j下标都++
		if haystack[i] == needle[j] { //i在for循环中++
			j++
		}

		//判断匹配结束 即下标j为模式串的长度，则当前文本串包含该模式串
		if j == len(needle) {
			return i - len(needle) + 1
		}
	}

	//文本串不包含模式串
	return -1
}

func main() {
	haystack := "aabaabaaf"
	needle := "aabaaf"
	flag := strStr(haystack, needle)
	fmt.Println(flag)
}
