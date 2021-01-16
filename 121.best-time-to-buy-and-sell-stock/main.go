package main

/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
func maxProfit(prices []int) int {
	min := prices[0]
	res := 0

	for i := 0; i < len(prices); i++ {
		p := prices[i]
		if p < min {
			min = p
		}
		if p-min > res {
			res = p - min
		}
	}

	return res
}

// @lc code=end
