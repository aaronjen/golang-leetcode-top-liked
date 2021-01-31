package main

/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	res := 0
	left := 0
	right := len(height) - 1

	for left < right {
		lh := height[left]
		rh := height[right]
		if lh > rh {
			res = max(res, rh*(right-left))
			right--
		} else {
			res = max(res, lh*(right-left))
			left++
		}
	}

	return res
}

// @lc code=end
