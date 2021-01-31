package main

/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}

	backtrack(candidates, target, &res, []int{}, 0)
	return res
}

func backtrack(candidates []int, target int, res *[][]int, tmp []int, start int) {
	if target < 0 {
		return
	}
	if target == 0 {
		tt := make([]int, len(tmp))
		copy(tt, tmp)
		*res = append(*res, tt)
	} else {
		for i := start; i < len(candidates); i++ {
			tmp = append(tmp, candidates[i])
			backtrack(candidates, target-candidates[i], res, tmp, i)
			tmp = tmp[:len(tmp)-1]
		}
	}
}

// @lc code=end
