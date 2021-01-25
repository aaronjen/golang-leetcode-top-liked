package main

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	table := map[int]int{}

	for i := 0; i < len(nums); i++ {
		if prev, ok := table[nums[i]]; ok {
			return []int{prev, i}
		}
		table[target-nums[i]] = i
	}

	return []int{}
}

// @lc code=end
