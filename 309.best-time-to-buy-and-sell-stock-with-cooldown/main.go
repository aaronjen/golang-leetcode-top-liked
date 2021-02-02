package main

import "math"

/*
 * @lc app=leetcode id=309 lang=golang
 *
 * [309] Best Time to Buy and Sell Stock with Cooldown
 */

// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	s1 := 0
	s2 := math.MinInt32
	s3 := 0
	for i := 0; i < len(prices); i++ {
		s1, s2, s3 = s3, max(s2, s1-prices[i]), max(s3, s2+prices[i])
	}
	return s3
}

// @lc code=end
