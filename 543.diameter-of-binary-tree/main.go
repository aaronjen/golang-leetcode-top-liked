package main

/*
 * @lc app=leetcode id=543 lang=golang
 *
 * [543] Diameter of Binary Tree
 */

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start

func diameterOfBinaryTree(root *TreeNode) int {
	res := 0
	dfs(root, &res)
	return res
}

func dfs(node *TreeNode, res *int) int {
	if node == nil {
		return 0
	}

	left := dfs(node.Left, res)
	right := dfs(node.Right, res)

	if left+right+1 > *res {
		*res = left + right
	}

	if left > right {
		return left + 1
	}
	return right + 1
}

// @lc code=end
