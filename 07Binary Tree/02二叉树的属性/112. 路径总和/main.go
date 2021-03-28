package main

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
1.递归函数的参数及返回值
root *TreeNode根节点
targetSum int

不用遍历完整棵树  有一条路径和为22即可返回true
*/

func traversal(root *TreeNode, targetSum int) bool {
	//2.确定递归终止条件
	//当到达叶子节点 且 targetSum减去路径上的数值最后为0时
	if root.Left == nil && root.Right == nil && targetSum == 0 {
		return true
	}

	//如果到达了叶子节点 但是targetSum还未减至0则说明该条路径无效
	if root.Left == nil && root.Right == nil {
		return false
	}

	if root.Left != nil {
		if traversal(root.Left, targetSum-root.Left.Val) {
			return true
		}
	}

	if root.Right != nil {
		if traversal(root.Right, targetSum-root.Right.Val) {
			return true
		}
	}

	return false
}

// func hasPathSum(root *TreeNode, targetSum int) bool {
// 	if root == nil {
// 		return false
// 	}
// 	return traversal(root, targetSum-root.Val)
// }

/*2*/
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil { //节点未null 之间返回false
		return false
	}

	//遍历到了叶子节点  此时看目标值减去当前节点的val是否为0 如果为零说明找到一条路径满足条件
	if root.Left == nil && root.Right == nil {
		return targetSum-root.Val == 0
	}

	//递归左右子树  只要有一个子树满足条件即可
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}
