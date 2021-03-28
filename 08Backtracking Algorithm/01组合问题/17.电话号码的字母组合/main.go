package main

import "fmt"

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
