package main

/*
 * @lc app=leetcode id=437 lang=golang
 *
 * [437] Path Sum III
 */

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
func pathSum(root *TreeNode, sum int) int {
	preSum := map[int]int{}
	preSum[0]++
	return helper(root, 0, sum, preSum)
}

func helper(node *TreeNode, curSum int, target int, preSum map[int]int) int {
	if node == nil {
		return 0
	}
	curSum += node.Val
	res := preSum[curSum-target]
	preSum[curSum]++
	res += helper(node.Left, curSum, target, preSum) + helper(node.Right, curSum, target, preSum)
	preSum[curSum]--
	return res
}

// @lc code=end
