package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//我们验证该二叉树是否对称，比较的左右节点，而是比较左右子树是否相同所以需要递归遍历两棵子树

/*
递归三部曲：
1.确定递归函数的参数和返回值：很明显，这里的参数是两个子树，返回值是bool类型
compare(left,right *TreeNode) bool

2.终止条件：有4种情况
左子树为null,右子树不为null return false
左子树不为null。右子树为null return falase
左子树和右子树都为null 说明对称 return true
左右子树都不为null 但是Val值不相同  return false

3.单层递归逻辑：
进入到了左右子树都不为null 并且Val值相同的情况了
此时比较外围的节点状况 left.Left 与 right.Right
内围情况 left.Right 与 right.Left
只有当内外围结果都相同的时候才对称
*/

func compare(left, right *TreeNode) bool {
	if left == nil && right != nil {
		return false
	} else if left != nil && right == nil {
		return false
	} else if left == nil && right == nil {
		return true
	} else if left.Val != right.Val {
		return false
	}

	//进入单层递归逻辑
	l := compare(left.Left, right.Right)
	r := compare(left.Right, right.Left)
	return l && r
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return compare(root.Left, root.Right)
}
