package main

import "sort"

/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 */

// @lc code=start
type item struct {
	val  int
	freq int
}

func topKFrequent(nums []int, k int) []int {
	table := map[int]int{}

	for i := 0; i < len(nums); i++ {
		table[nums[i]]++
	}

	ar := []item{}
	for key, freq := range table {
		ar = append(ar, item{
			val:  key,
			freq: freq,
		})
	}

	sort.Slice(ar, func(i, j int) bool {
		return ar[i].freq > ar[j].freq
	})

	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = ar[i].val
	}
	return res
}

// @lc code=end
