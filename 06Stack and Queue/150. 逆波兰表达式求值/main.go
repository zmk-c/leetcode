package main

import (
	"fmt"
	"strconv"
)

//遍历逆波兰表达式，遇到数字进栈，当遇到数学运算符时，
//将栈里的前两个元素分别出栈再与运算符运算，得到的数再入栈。
func evalRPN(tokens []string) int {
	stack := make([]int, 0)

	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "+" || tokens[i] == "-" || tokens[i] == "*" || tokens[i] == "/" {
			//将栈顶前两个元素出栈
			num1 := stack[len(stack)-1]  //获取栈顶元素
			stack = stack[:len(stack)-1] //pop出栈 更新栈
			num2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			//判断各个运算符 进行对应得运算操作后再push进栈
			if tokens[i] == "+" {
				stack = append(stack, num2+num1) //注意运算顺序
			}
			if tokens[i] == "-" {
				stack = append(stack, num2-num1)
			}
			if tokens[i] == "*" {
				stack = append(stack, num2*num1)
			}
			if tokens[i] == "/" {
				stack = append(stack, num2/num1)
			}
		} else { //说明遇到得时数字  string -> int入栈
			num, _ := strconv.Atoi(tokens[i])
			stack = append(stack, num)
		}
	}
	result := stack[len(stack)-1]
	return result
}

func main() {
	tokens := []string{"2", "1", "+", "3", "*"}
	result := evalRPN(tokens)
	fmt.Println(result)
}
