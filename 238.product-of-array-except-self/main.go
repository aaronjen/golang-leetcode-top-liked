package main

/*
 * @lc app=leetcode id=238 lang=golang
 *
 * [238] Product of Array Except Self
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = 1

	prev := 1
	for i := 1; i < len(nums); i++ {
		prev *= nums[i-1]
		res[i] = prev
	}

	prev = 1
	for i := len(nums) - 2; i >= 0; i-- {
		prev *= nums[i+1]
		res[i] *= prev
	}

	return res
}

// @lc code=end
