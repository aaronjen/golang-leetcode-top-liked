package main

/*
 * @lc app=leetcode id=494 lang=golang
 *
 * [494] Target Sum
 */

// @lc code=start
func findTargetSumWays(nums []int, S int) int {
	if S > 1000 || S < -1000 {
		return 0
	}
	dp := make([]int, 2001)
	dp[1000] = 1
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		tmp := make([]int, 2001)
		for j := 0; j < 2001; j++ {
			if dp[j] > 0 {
				if j+n < 2001 {
					tmp[j+n] += dp[j]
				}
				if j-n >= 0 {
					tmp[j-n] += dp[j]
				}
			}
		}
		dp = tmp
	}
	return dp[S+1000]
}

// @lc code=end
