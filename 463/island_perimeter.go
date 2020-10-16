// 463. Island Perimeter
// https://leetcode.com/problems/island-perimeter/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(islandPerimeter([][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}))
	pp.Println(16)
	pp.Println("=========================================")
	pp.Println(islandPerimeter([][]int{
		{1, 0},
	}))
	pp.Println(4)
	pp.Println("=========================================")
	pp.Println(islandPerimeter([][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}))
	pp.Println(16)
	pp.Println("=========================================")
	pp.Println(islandPerimeter([][]int{
		{1, 1},
		{1, 1},
	}))
	pp.Println(8)
	pp.Println("=========================================")
}

// 1 Runtime Error
// 2 Wrong Answer
// 40 mins

func islandPerimeter(grid [][]int) int {
	var x, y int
FIND_INITIAL_POINT:
	for y = range grid {
		for x = range grid[y] {
			if grid[y][x] == 1 {
				break FIND_INITIAL_POINT
			}
		}
	}

	// check := func(x, y int) int {
	// 	if x < 0 || x >= len(grid[0]) || y < 0 || y >= len(grid) || grid[y][x] == 0 {
	// 		return 1
	// 	}
	// 	return -1
	// }

	// var dfs func(x, y int) int
	// dfs = func(x, y int) int {
	// 	if x < 0 || x >= len(grid[0]) || y < 0 || y >= len(grid) || grid[y][x] == 0 {
	// 		return 0
	// 	}
	// 	grid[y][x] = 0
	// 	return check(x, y-1) + check(x+1, y) + check(x, y+1) + check(x-1, y) + dfs(x, y-1) + dfs(x+1, y) + dfs(x, y+1) + dfs(x-1, y)
	// }

	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		if x < 0 || x >= len(grid[0]) || y < 0 || y >= len(grid) || grid[y][x] == 0 {
			return 1
		}
		grid[y][x] = 0
		return dfs(x, y-1) + dfs(x+1, y) + dfs(x, y+1) + dfs(x-1, y) - 1
	}

	return dfs(x, y) + 1
}
