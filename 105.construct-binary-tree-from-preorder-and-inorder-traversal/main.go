package main

/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
 */

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start

func buildTree(preorder []int, inorder []int) *TreeNode {
	ind := 0
	return build(preorder, inorder, &ind, 0, len(inorder))
}

func build(preorder []int, inorder []int, ind *int, start, end int) *TreeNode {
	if *ind >= len(preorder) || start > end {
		return nil
	}
	val := preorder[*ind]
	pos := start
	for val != inorder[pos] {
		pos++
	}
	*ind++
	return &TreeNode{
		Val:   val,
		Left:  build(preorder, inorder, ind, start, pos-1),
		Right: build(preorder, inorder, ind, pos+1, end),
	}
}

// @lc code=end
