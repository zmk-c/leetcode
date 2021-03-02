package main

import "fmt"

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。

示例 1：
输入：s = "()"
输出：true

示例 2：
输入：s = "()[]{}"
输出：true

示例 3：
输入：s = "(]"
输出：false

示例 4：
输入：s = "([)]"
输出：false

示例 5：
输入：s = "{[]}"
输出：true

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func isValid(s string) bool {
	//用栈来做匹配问题是非常方便的

	//遍历字符串s 如果发现左括号  则将对应得右括号push进栈  减少了空间复杂度
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, ')')
		} else if s[i] == '[' {
			stack = append(stack, ']')
		} else if s[i] == '{' {
			stack = append(stack, '}')
		} else if len(stack) == 0 || stack[len(stack)-1] != s[i] { //开始匹配  第三种情况和第二种情况
			return false
		} else {
			stack = stack[:len(stack)-1] //说明栈顶元素与对应右括号匹配 pop出栈
		}
	}
	if len(stack) != 0 { //遍历完整个字符串s后如果当前栈不为空  则说明第一种情况
		return false
	}
	return true //所有括号都匹配
}

func main() {
	s := "([{}])"
	flag := isValid(s)
	fmt.Println(flag)
}
