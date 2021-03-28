package main

import (
	"fmt"
	"strings"
)

/*思路：
1.移除多余的空格
2.将每个单词反转*/
func reverseWords(s string) string {
	newS := removeExtraSpace(s)
	words := strings.Split(newS, " ")

	l := 0
	r := len(words) - 1

	for l < r {
		words[l], words[r] = words[r], words[l]
		l++
		r--
	}

	return strings.Join(words, " ")
}

func removeExtraSpace(s string) string {
	//用一个切片字节数组保存移除多余空格后的字符
	var result []byte
	//字符串长度
	l := len(s)
	//用于处理连续多个空格
	flag := true

	for l != 0 && s[l-1] == ' ' { //移除字符串后面的空格
		l--
	}

	for i := 0; i < l; i++ {
		if s[i] != ' ' { //如果字符不为空格  表示遇到了单词字符  需要加入新的数组
			result = append(result, s[i])
			flag = false
		}
		if s[i] == ' ' && flag == false { //①排除字符串开始的空格
			result = append(result, s[i])
			flag = true //②排除字符串中间连续多个空格
		}
	}
	return string(result)
}

func main() {
	s := "  hello world!  "
	newS := reverseWords(s)
	fmt.Println(newS)
}
