package main

import "fmt"

/*
给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
返回 s 所有可能的分割方案。

示例:

输入: "aab"
输出:
[
  ["aa","b"],
  ["a","a","b"]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-partitioning
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
切割问题类似于组合问题
回溯三部曲：
1.参数的确定
2.终止条件
3.单层循环实现
*/

/*
参数说明：
s string 题目所给的需要切割的字符串
startIndex int 切割的起始位置，即作为切割线
path []string 存放回文的子串，一组解
result *[][]string 最终解集合
*/
func backtrack(s string, startIndex int, path *[]string, result *[][]string) {
	//终止条件：切割线切割到了字符串的最后面  说明找到了一个切割方法
	if startIndex >= len(s) {
		tmp := make([]string, len(*path))
		copy(tmp, *path)
		*result = append(*result, tmp)
		return
	}

	//单层循环实现
	for i := startIndex; i < len(s); i++ {
		//判断是不是回文串，是的话 加入path,否的话跳过
		if !isPalindrome(s, startIndex, i) { //注意截取的位置
			continue
		}
		subStr := s[startIndex : i+1]
		*path = append(*path, subStr)
		backtrack(s, i+1, path, result) //切割过的位置不能重复切割
		*path = (*path)[:len(*path)-1]
	}
}

//判断该字符串是否是回文串
func isPalindrome(s string, start, end int) bool {
	//利用双指针遍历字符串头尾比较
	for i, j := start, end; i < j; i, j = i+1, j-1 { //for循环中多个参数变量形式注意一下写法
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func partition(s string) [][]string {
	path := make([]string, 0)
	result := make([][]string, 0)
	if len(s) == 0 {
		return result
	}
	backtrack(s, 0, &path, &result)
	return result
}

func main() {
	s := "aab"
	result := partition(s)
	fmt.Println(result)
}
