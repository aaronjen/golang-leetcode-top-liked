package main

/*
 * @lc app=leetcode id=101 lang=golang
 *
 * [101] Symmetric Tree
 */

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	root.Left = rotate(root.Left)

	return same(root.Left, root.Right)
}

func rotate(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	node.Left, node.Right = rotate(node.Right), rotate(node.Left)

	return node
}

func same(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}

	if t1 == nil || t2 == nil {
		return false
	}

	if t1.Val != t2.Val {
		return false
	}

	return same(t1.Left, t2.Left) && same(t1.Right, t2.Right)
}

// @lc code=end
