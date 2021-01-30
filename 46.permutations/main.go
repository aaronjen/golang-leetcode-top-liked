package main

/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start
func permute(nums []int) [][]int {
	res := [][]int{}
	used := make([]bool, len(nums))
	backtrack(&res, nums, []int{}, used)

	return res
}

func backtrack(res *[][]int, nums []int, tmp []int, used []bool) {
	if len(tmp) == len(nums) {
		tt := make([]int, len(tmp))
		copy(tt, tmp)
		*res = append(*res, tt)
	} else {
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			tmp = append(tmp, nums[i])
			backtrack(res, nums, tmp, used)
			tmp = tmp[:len(tmp)-1]
			used[i] = false
		}
	}
}

// @lc code=end
