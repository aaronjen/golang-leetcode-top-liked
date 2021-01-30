package main

/*
 * @lc app=leetcode id=763 lang=golang
 *
 * [763] Partition Labels
 */

// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func partitionLabels(S string) []int {
	table := map[byte]int{}

	for i := 0; i < len(S); i++ {
		b := S[i]
		table[b] = i
	}

	res := []int{}
	start := -1
	curMax := 0
	for i := 0; i < len(S); i++ {
		b := S[i]
		curMax = max(curMax, table[b])
		if i == curMax {
			res = append(res, curMax-start)
			start = curMax
		}
	}
	return res
}

// @lc code=end
