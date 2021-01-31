package main

import "sort"

/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	res := [][]string{}

	table := map[string][]string{}

	for i := 0; i < len(strs); i++ {
		bs := []byte(strs[i])
		sort.Slice(bs, func(i, j int) bool {
			return bs[i] < bs[j]
		})
		s := string(bs)
		if table[s] == nil {
			table[s] = []string{strs[i]}
		} else {
			table[s] = append(table[s], strs[i])
		}
	}

	for _, v := range table {
		res = append(res, v)
	}

	return res
}

// @lc code=end
