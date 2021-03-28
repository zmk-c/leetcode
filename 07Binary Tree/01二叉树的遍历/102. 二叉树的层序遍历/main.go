package main

import "fmt"

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//层次遍历使用队列来作为辅助
// func levelOrder(root *TreeNode) [][]int {
// 	//辅助队列 用于存放节点
// 	var queue []*TreeNode
// 	//最终结果  二维切片中存放每一层的节点值
// 	var result [][]int
// 	if root != nil {
// 		queue = append(queue, root)
// 	}
// 	//如过队列不为空，则从队列中获取当前出口处的节点
// 	for len(queue) != 0 {
// 		//存放每一层节点的值
// 		var res []int
// 		size := len(queue)
// 		for i := 0; i < size; i++ {
// 			node := queue[0]
// 			res = append(res, node.Val)
// 			queue = queue[1:]

// 			//继续向队列中追加当前节点的左右子树
// 			if node.Left != nil {
// 				queue = append(queue, node.Left)
// 			}

// 			if node.Right != nil {
// 				queue = append(queue, node.Right)
// 			}
// 		}
// 		result = append(result, res)
// 	}
// 	return result
// }

//dfs递归操作
var result [][]int

func levelOrder(root *TreeNode) [][]int {
	result = [][]int{}
	dfs(root, 0)
	return result
}

func dfs(root *TreeNode, step int) {
	//不是每次进入dfs都需要重新创建一个空的tmp  如果tmp中有数据就不需要新创建了
	if root != nil {
		if step == len(result) {
			result = append(result, []int{})
		}
		result[step] = append(result[step], root.Val)
		dfs(root.Left, step+1)
		dfs(root.Right, step+1)
	}
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	result := levelOrder(root)
	fmt.Println(result)
}
