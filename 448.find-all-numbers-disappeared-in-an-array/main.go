package main

/*
 * @lc app=leetcode id=448 lang=golang
 *
 * [448] Find All Numbers Disappeared in an Array
 */

// @lc code=start
func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		for nums[n-1] != n {
			nums[n-1], nums[i] = nums[i], nums[n-1]
			n = nums[i]
		}
	}

	res := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			res = append(res, i+1)
		}
	}

	return res
}

// @lc code=end
