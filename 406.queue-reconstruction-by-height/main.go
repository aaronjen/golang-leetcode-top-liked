package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=406 lang=golang
 *
 * [406] Queue Reconstruction by Height
 */

// @lc code=start
func insert(q [][]int, pos int, p []int) {
	copy(q[pos+1:], q[pos:])
	q[pos] = p
}

func reconstructQueue(people [][]int) [][]int {
	res := make([][]int, len(people))
	for i := 0; i < len(people); i++ {
		res[i] = make([]int, 2)
	}

	sort.Slice(people, func(i, j int) bool {
		a := people[i]
		b := people[j]
		if a[0] != b[0] {
			return a[0] > b[0]
		}
		return a[1] < b[1]
	})

	for _, p := range people {
		insert(res, p[1], p)
	}

	return res
}

// @lc code=end
