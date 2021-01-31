package main

/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */

// @lc code=start
func subsets(nums []int) [][]int {
	res := [][]int{}

	backtrack(&res, nums, []int{}, 0)

	return res
}

func backtrack(res *[][]int, nums []int, tmp []int, start int) {
	tt := make([]int, len(tmp))
	copy(tt, tmp)
	*res = append(*res, tt)

	for i := start; i < len(nums); i++ {
		tmp = append(tmp, nums[i])
		backtrack(res, nums, tmp, i+1)
		tmp = tmp[:len(tmp)-1]
	}
}

// @lc code=end
