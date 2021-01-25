package main

/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

// @lc code=start
func climbStairs(n int) int {
	prev1 := 1
	prev2 := 1

	for i := 2; i <= n; i++ {
		prev1, prev2 = prev2, prev1+prev2
	}

	return prev2
}

// @lc code=end
