package main

/*
 * @lc app=leetcode id=394 lang=golang
 *
 * [394] Decode String
 */

// @lc code=start
func isNum(b byte) bool {
	return b >= '0' && b <= '9'
}

func decodeString(s string) string {
	n := 0
	sb := ""
	nStack := []int{}
	sStack := []string{}
	for i := 0; i < len(s); i++ {
		b := s[i]
		if isNum(b) {
			n = n*10 + int(b-'0')
		} else if b == '[' {
			nStack = append(nStack, n)
			sStack = append(sStack, sb)
			n = 0
			sb = ""
		} else if b == ']' {
			tmpN := nStack[len(nStack)-1]
			nStack = nStack[:len(nStack)-1]
			prevS := sStack[len(sStack)-1]
			sStack = sStack[:len(sStack)-1]
			tmpS := ""
			for i := 0; i < tmpN; i++ {
				tmpS += sb
			}
			sb = prevS + tmpS
		} else {
			sb += string(b)
		}
	}

	return sb
}

// @lc code=end
