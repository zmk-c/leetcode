package main

import "fmt"

//给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

/*
回溯三部曲：
	1.递归函数参数返回值
	2.确定终止条件
	3.单层递归逻辑
*/

/*
参数说明：
n		1~n 个数
k		k位的组合
startIndex	开始遍历的位置 默认从1开始到n
path	访问的路径组合列表
result	所以k长度组合的结果集
*/

//单纯的回溯  未使用剪枝函数
func backtrack(n, k, startIndex int, path []int, result *[][]int) {
	//确定终止条件：当前路径长度=k
	if k == len(path) {
		//将当前组合结果加入总结果集
		tmp := make([]int, k)
		copy(tmp, path)
		*result = append(*result, tmp)
		return
	}

	//单层递归逻辑
	/*
		组合问题剪枝函数的设计：
			已经选择的元素：len(path)
			还需要选择的元素：k-len(path)
			在集合n中至多要从该起始位置：n-(k-len(path))+1
	*/
	for i := startIndex; i <= n-(k-len(path))+1; i++ { //原来是i<=n  这里改为 i<=n-(k-len(path))+1 剪枝函数
		//向路径中添加该元素
		path = append(path, i)
		//递归 向下一层遍历，起点位置+1
		backtrack(n, k, i+1, path, result)
		//回溯
		path = path[:len(path)-1]
	}
}

func combine(n int, k int) [][]int {
	path := make([]int, 0)
	result := make([][]int, 0)
	backtrack(n, k, 1, path, &result)
	return result
}

func main() {
	n := 4
	k := 2
	result := combine(n, k)
	fmt.Printf("1~%v的长度为%v的组合共有%v种:\n分别是%v", n, k, len(result), result)
}
