package main

/*
 * @lc app=leetcode id=647 lang=golang
 *
 * [647] Palindromic Substrings
 */

// @lc code=start
func countSubstrings(s string) int {
	res := 0

	for i := 0; i < len(s); i++ {
		extend(s, i, i, &res)
		extend(s, i, i+1, &res)
	}

	return res
}

func extend(s string, left, right int, res *int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		*res++
		left--
		right++
	}
}

// @lc code=end
