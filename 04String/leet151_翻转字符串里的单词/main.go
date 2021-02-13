package main

import (
	"fmt"
	"strings"
)

/*
给定一个字符串，逐个翻转字符串中的每个单词。

说明：

无空格字符构成一个 单词 。
输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。

示例 1：
输入："the sky is blue"
输出："blue is sky the"

示例 2：
输入："  hello world!  "
输出："world! hello"
解释：输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。

示例 3：
输入："a good   example"
输出："example good a"
解释：如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-words-in-a-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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
