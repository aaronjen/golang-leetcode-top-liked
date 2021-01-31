package main

/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */

// @lc code=start
var table [][]byte = [][]byte{
	[]byte("abc"),
	[]byte("def"),
	[]byte("ghi"),
	[]byte("jkl"),
	[]byte("mno"),
	[]byte("pqrs"),
	[]byte("tuv"),
	[]byte("wxyz"),
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	res := []string{}

	backtrack(digits, 0, &res, []byte{})

	return res
}

func backtrack(digits string, ind int, res *[]string, tmp []byte) {
	if ind == len(digits) {
		*res = append(*res, string(tmp))
	} else {
		br := table[int(digits[ind]-'2')]
		for _, b := range br {
			tmp = append(tmp, b)
			backtrack(digits, ind+1, res, tmp)
			tmp = tmp[:len(tmp)-1]
		}
	}
}

// @lc code=end
