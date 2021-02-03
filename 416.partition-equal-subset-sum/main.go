package main

/*
 * @lc app=leetcode id=416 lang=golang
 *
 * [416] Partition Equal Subset Sum
 */

// @lc code=start
func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}
	sum /= 2
	dp := make([]bool, sum+1)
	dp[0] = true
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		for j := sum; j >= n; j-- {
			dp[j] = dp[j] || dp[j-n]
		}
	}

	return dp[sum]
}

// @lc code=end
