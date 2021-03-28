package main

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//最小深度是指从根节点到叶子节点的最小值  特别注意存在单棵子树为空的情况
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHight := minDepth(root.Left)
	rightHight := minDepth(root.Right)

	//如果左子树为null 则二叉树的最小深度为右子树的最小深度+1
	if root.Left == nil && root.Right != nil {
		return rightHight + 1
	}

	if root.Right == nil && root.Left != nil {
		return leftHight + 1
	}

	//比较左右子树的最小值
	if leftHight <= rightHight {
		return leftHight + 1
	} else {
		return rightHight + 1
	}
}
