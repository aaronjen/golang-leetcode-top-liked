package main

/*
 * @lc app=leetcode id=617 lang=golang
 *
 * [617] Merge Two Binary Trees
 */

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}

	val := 0
	var t1Left *TreeNode
	var t1Right *TreeNode
	var t2Left *TreeNode
	var t2Right *TreeNode

	if t1 != nil {
		val += t1.Val
		t1Left = t1.Left
		t1Right = t1.Right
	}
	if t2 != nil {
		val += t2.Val
		t2Left = t2.Left
		t2Right = t2.Right
	}

	return &TreeNode{
		Val:   val,
		Left:  mergeTrees(t1Left, t2Left),
		Right: mergeTrees(t1Right, t2Right),
	}
}

// @lc code=end
