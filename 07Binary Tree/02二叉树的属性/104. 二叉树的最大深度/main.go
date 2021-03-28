package main

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//遍历左右子树高度的最大值+1即为整棵二叉树的最大值
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHight := maxDepth(root.Left)
	rightHight := maxDepth(root.Right)

	if leftHight >= rightHight {
		return leftHight + 1
	} else {
		return rightHight + 1
	}
}
