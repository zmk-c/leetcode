package main

import (
	"fmt"
)

/*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

示例:

输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
说明:
尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*回溯三部曲：
1.确定回溯函数参数
2.终止条件
3.单层循环遍历

以及需要考虑的问题：
Q：数字和字符之间如何映射？
A：可以利用map或者二维数组做映射

Q:输入0 1 * #等异常情况
*/

var letters = [10]string{
	"",     //0
	"",     //1
	"abc",  //2
	"def",  //3
	"ghi",  //4
	"jkl",  //5
	"mno",  //6
	"pqrs", //7
	"tuv",  //8
	"wxyz", //9
}

/*参数解释：
digits string 题目所给的数组
letters []string 电话数字对应的字母
s string 路径上组成的字母字符串形成一个解
index int 标记digits中的数字
result *[]string 存放最终解集
*/
func backtrack(digits string, s string, index int, result *[]string) {
	//终止条件
	if index == len(digits) {
		//将该解加入result结果集合
		*result = append(*result, s)
		return
	}

	//将digits数字字符串 去对应的下标转换为letter中的字母字符串
	digit := digits[index] - '0'
	//fmt.Printf("%T,%v \n", digit, digit)
	letter := letters[digit]
	//fmt.Printf("%T,%v \n", letter, letter)
	for i := 0; i < len(letter); i++ {
		s += string(letter[i])
		backtrack(digits, s, index+1, result)
		s = s[:len(s)-1]
	}
}

func letterCombinations(digits string) []string {
	var s string
	result := make([]string, 0)
	if len(digits) == 0 {
		return result
	}
	backtrack(digits, s, 0, &result)
	return result
}

func main() {
	digits := "23"
	result := letterCombinations(digits)
	fmt.Println(result)
}
