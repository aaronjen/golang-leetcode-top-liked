package main

/*
 * @lc app=leetcode id=114 lang=golang
 *
 * [114] Flatten Binary Tree to Linked List
 */

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start

func flatten(root *TreeNode) {
	stack := []*TreeNode{}
	nd := root
	for nd != nil {
		if nd.Right != nil {
			stack = append(stack, nd.Right)
		}
		if nd.Left != nil {
			nd.Right = nd.Left
			nd.Left = nil
			nd = nd.Right
		} else if len(stack) > 0 {
			nd.Right = stack[len(stack)-1]
			nd = nd.Right
			stack = stack[:len(stack)-1]
		} else {
			break
		}
	}
}

// @lc code=end
