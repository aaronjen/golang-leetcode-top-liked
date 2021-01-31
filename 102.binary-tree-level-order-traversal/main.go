package main

/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
 */

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	var cur []int
	res := [][]int{}
	for len(queue) != 0 {
		cur = make([]int, len(queue))
		tmp := []*TreeNode{}
		for i, nd := range queue {
			cur[i] = nd.Val
			if nd.Left != nil {
				tmp = append(tmp, nd.Left)
			}
			if nd.Right != nil {
				tmp = append(tmp, nd.Right)
			}
		}
		res = append(res, cur)
		queue = tmp
	}

	return res
}

// @lc code=end
