package main

/*
 * @lc app=leetcode id=739 lang=golang
 *
 * [739] Daily Temperatures
 */

// @lc code=start
func dailyTemperatures(T []int) []int {
	res := make([]int, len(T))

	for i := len(T) - 2; i >= 0; i-- {
		future := i + 1
		for future < len(T) {
			if T[future] > T[i] {
				res[i] = future - i
				break
			}
			if res[future] > 0 {
				future += res[future]
			} else {
				break
			}
		}
	}

	return res
}

// @lc code=end
