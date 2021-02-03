package main

/*
 * @lc app=leetcode id=438 lang=golang
 *
 * [438] Find All Anagrams in a String
 */

// @lc code=start
func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	need := make([]int, 26)
	cnt := 0
	n := len(p)
	for i := 0; i < n; i++ {
		pos := int(p[i] - 'a')
		need[pos]++
		if need[pos] == 1 {
			cnt++
		}
	}
	for i := 0; i < n; i++ {
		pos := int(s[i] - 'a')
		need[pos]--
		if need[pos] == 0 {
			cnt--
		} else if need[pos] == -1 {
			cnt++
		}
	}

	res := []int{}
	for i := n; i < len(s); i++ {
		if cnt == 0 {
			res = append(res, i-n)
		}
		pos1, pos2 := int(s[i-n]-'a'), int(s[i]-'a')
		need[pos1]++
		if need[pos1] == 0 {
			cnt--
		} else if need[pos1] == 1 {
			cnt++
		}

		need[pos2]--
		if need[pos2] == 0 {
			cnt--
		} else if need[pos2] == -1 {
			cnt++
		}
	}

	if cnt == 0 {
		res = append(res, len(s)-n)
	}
	return res
}

// @lc code=end
