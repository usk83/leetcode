// 200. Number of Islands
// https://leetcode.com/problems/number-of-islands/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(numIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(numIslands([][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}))
	pp.Println(3)
	pp.Println("=========================================")
}

func numIslands(grid [][]byte) int {
	var dfs func(int, int)
	dfs = func(row, col int) {
		if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[row]) || grid[row][col] == '0' {
			return
		}
		grid[row][col] = '0'
		dfs(row-1, col)
		dfs(row, col+1)
		dfs(row+1, col)
		dfs(row, col-1)
	}

	var count int
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == '1' {
				count++
				dfs(row, col)
			}
		}
	}

	return count
}
