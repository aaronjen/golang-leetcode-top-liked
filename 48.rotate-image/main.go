package main

/*
 * @lc app=leetcode id=48 lang=golang
 *
 * [48] Rotate Image
 */

// @lc code=start

func rotate(matrix [][]int) {
	swap := func(row1, col1, row2, col2 int) {
		matrix[row1][col1], matrix[row2][col2] = matrix[row2][col2], matrix[row1][col1]
	}

	m := len(matrix)
	n := len(matrix[0])

	for i := 0; i < m/2; i++ {
		for j := 0; j < n; j++ {
			swap(i, j, m-i-1, j)
		}
	}

	for i := 0; i < m; i++ {
		for j := i + 1; j < n; j++ {
			swap(i, j, j, i)
		}
	}
}

// @lc code=end
