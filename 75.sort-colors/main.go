package main

/*
 * @lc app=leetcode id=75 lang=golang
 *
 * [75] Sort Colors
 */

// @lc code=start
func sortColors(nums []int) {
	left := -1
	right := len(nums)
	for i := 0; i < len(nums); i++ {
		for nums[i] != 1 && i > left && i < right {
			if nums[i] == 2 {
				right--
				nums[i], nums[right] = nums[right], nums[i]
			} else {
				left++
				nums[i], nums[left] = nums[left], nums[i]
			}
		}
	}
}

// @lc code=end
