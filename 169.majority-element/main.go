package main

/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 */

// @lc code=start
func majorityElement(nums []int) int {
	res := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == res {
			count++
		} else {
			count--
			if count == -1 {
				res = nums[i]
				count = 1
			}
		}
	}

	return res
}

// @lc code=end
