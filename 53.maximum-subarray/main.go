package main

/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
func findmax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxSubArray(nums []int) int {
	res := nums[0]
	cur := nums[0]
	for i := 1; i < len(nums); i++ {
		if cur > 0 {
			cur += nums[i]
		} else {
			cur = nums[i]
		}

		res = findmax(res, cur)
	}

	return res
}

// @lc code=end
