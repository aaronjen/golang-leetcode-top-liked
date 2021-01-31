package main

/*
 * @lc app=leetcode id=287 lang=golang
 *
 * [287] Find the Duplicate Number
 */

// @lc code=start
func findDuplicate(nums []int) int {
	fast := nums[nums[0]]
	slow := nums[0]
	for fast != slow {
		fast = nums[nums[fast]]
		slow = nums[slow]
	}
	fast = 0
	for fast != slow {
		fast = nums[fast]
		slow = nums[slow]
	}

	return fast
}

// @lc code=end
