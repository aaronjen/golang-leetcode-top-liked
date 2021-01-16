package main

/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 */

// @lc code=start
func moveZeroes(nums []int) {
	count := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[count], nums[i] = nums[i], nums[count]
			count++
		}
	}
}

// @lc code=end
