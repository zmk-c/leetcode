package main

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*递归三部曲
1.确定递归函数的参数和返回值
2.确定终止条件
3.确定单层递归的逻辑
*/

//1.确定递归函数的参数和返回值：因为要打印出后序遍历节点的数值，所以参数里传入[]int类型
//不需要有返回值
// func traversal(root *TreeNode, result *[]int) {
// 	//2.确定终止条件
// 	if root == nil {
// 		return
// 	}

// 	//3.确定单层递归的逻辑
// 	//后序遍历 是 左 右 中 的顺序，所以先取左节点的值
// 	traversal(root.Left, result)        //左
// 	traversal(root.Right, result)       //右
// 	*result = append(*result, root.Val) //中
// }

// func postorderTraversal(root *TreeNode) []int {
// 	result := make([]int, 0)
// 	traversal(root, &result)
// 	return result
// }

//非递归
//后序遍历的顺序  左 右 中
//前序遍历的顺序 为中 左 右  -> 更改出栈顺序为->中 右 左 -> 反转数组 -> 左 右 中 的到后序遍历
func postorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0) //存放二叉树节点
	result := make([]int, 0)      //存放节点数值
	stack = append(stack, root)   //首先将根节点入栈
	for len(stack) != 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur != nil {
			result = append(result, cur.Val) //中
		} else {
			continue //排除叶子节点所指向的左右孩子为空
		}

		stack = append(stack, cur.Left)  //左子树进栈
		stack = append(stack, cur.Right) //右子树进栈
	}
	//则整个result数组顺序为 二叉树的 中 右 左  此时反转result数组就得到了后序遍历的结果
	for i, j := 0, len(result)-1; i < j; i++ {
		result[i], result[j] = result[j], result[i]
		j--
	}

	return result
}
