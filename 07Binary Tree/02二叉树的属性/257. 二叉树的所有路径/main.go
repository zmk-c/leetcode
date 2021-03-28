package main

import (
	"fmt"
	"strconv"
)

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
1.参数及其返回值：
root *TreeNode 根节点
path string 用来存放遍历的路径
result *[]string 存放最终的路径结果集
*/
//遍历右子树之前，路径要撤销最末尾的选择，如果path用的是数组，就会弹出最后一项。
//这里用的字符串，path保存了当前节点的路径，递归右子树时，传入它即可，它不包含在递归左子树所拼接的东西。
func travesal(root *TreeNode, path string, result *[]string) {
	//2.递归终止条件 到达叶子节点
	if root.Left == nil && root.Right == nil {
		path += strconv.Itoa(root.Val)
		//加入结果集中
		*result = append(*result, path)
		return
	}

	//处理非叶子节点
	path += strconv.Itoa(root.Val) + "->"
	if root.Left != nil {
		travesal(root.Left, path, result)
	}
	if root.Right != nil {
		travesal(root.Right, path, result)
	}
}

func binaryTreePaths(root *TreeNode) []string {
	result := []string{}
	travesal(root, "", &result)
	return result
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	result := binaryTreePaths(root)
	fmt.Println(result)
}
