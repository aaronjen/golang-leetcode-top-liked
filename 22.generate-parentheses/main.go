package main

/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */

// @lc code=start
func generateParenthesis(n int) []string {
	res := []string{}

	backtrack(&res, n, 0, []byte{})

	return res
}

func backtrack(res *[]string, n, open int, tmp []byte) {
	if n == 0 && open == 0 {
		*res = append(*res, string(tmp))
	} else {
		if n > 0 {
			tmp = append(tmp, '(')
			backtrack(res, n-1, open+1, tmp)
			tmp = tmp[:len(tmp)-1]
		}
		if open > 0 {
			tmp = append(tmp, ')')
			backtrack(res, n, open-1, tmp)
			tmp = tmp[:len(tmp)-1]
		}
	}
}

// @lc code=end
