package main

/*
 * @lc app=leetcode id=621 lang=golang
 *
 * [621] Task Scheduler
 */

// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}

	counts := make([]int, 26)
	maxCount := 0
	maxN := 0
	for i := 0; i < len(tasks); i++ {
		t := tasks[i]
		pos := int(t - 'A')
		counts[pos]++
		if counts[pos] > maxN {
			maxN = counts[pos]
			maxCount = 1
		} else if counts[pos] == maxN {
			maxCount++
		}
	}

	return max((n+1)*(maxN-1)+maxCount, len(tasks))
}

// @lc code=end

// "["A","A","A","A","A","A","B","C","D","E","F","G"]\n1"
