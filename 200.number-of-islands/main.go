package main

/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res++
				mark(grid, i, j)
			}
		}
	}
	return res
}

func mark(grid [][]byte, row, col int) {
	m := len(grid)
	n := len(grid[0])

	if row < 0 || row >= m || col < 0 || col >= n || grid[row][col] == '0' {
		return
	}
	grid[row][col] = '0'
	mark(grid, row-1, col)
	mark(grid, row+1, col)
	mark(grid, row, col-1)
	mark(grid, row, col+1)
}

// @lc code=end
