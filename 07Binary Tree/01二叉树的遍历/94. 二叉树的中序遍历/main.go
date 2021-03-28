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

//1.确定递归函数的参数和返回值：因为要打印出中序遍历节点的数值，所以参数里传入[]int类型
//不需要有返回值
// func traversal(root *TreeNode, result *[]int) {
// 	//2.确定终止条件
// 	if root == nil {
// 		return
// 	}

// 	//3.确定单层递归的逻辑
// 	//中序遍历 是 左 中 右 的顺序
// 	traversal(root.Left, result)        //左
// 	*result = append(*result, root.Val) //中
// 	traversal(root.Right, result)       //右
// }

// func inorderTraversal(root *TreeNode) []int {
// 	result := make([]int, 0)
// 	traversal(root, &result)
// 	return result
// }

//非递归
func inorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0) //存放二叉树节点
	result := make([]int, 0)      //存放节点数值

	for root != nil || len(stack) != 0 {
		if root != nil {
			//将当前节点存入栈
			stack = append(stack, root)
			root = root.Left //左  一直访问到最左边的节点
		} else { //说明到达了二叉树的最底部
			//从栈中取出当前节点
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			//遍历节点的数值
			result = append(result, root.Val) //中
			root = root.Right                 //右
		}
	}

	return result
}
