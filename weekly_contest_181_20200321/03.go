// 5366. Check if There is a Valid Path in a Grid
// https://leetcode.com/contest/weekly-contest-181/problems/check-if-there-is-a-valid-path-in-a-grid/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(hasValidPath([][]int{{2, 4, 3}, {6, 5, 2}}))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(hasValidPath([][]int{{1, 2, 1}, {1, 2, 1}}))
	pp.Println(false)
	pp.Println("=========================================")
	pp.Println(hasValidPath([][]int{{1, 1, 2}}))
	pp.Println(false)
	pp.Println("=========================================")
	pp.Println(hasValidPath([][]int{{1, 1, 1, 1, 1, 1, 3}}))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(hasValidPath([][]int{{2}, {2}, {2}, {2}, {2}, {2}, {6}}))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(hasValidPath([][]int{{4, 1}, {6, 1}}))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(hasValidPath([][]int{{6, 1, 3}, {4, 1, 5}}))
	pp.Println(true)
	pp.Println("=========================================")
}

// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 300
// 1 <= grid[i][j] <= 6
func hasValidPath(grid [][]int) bool {
	y, x := 0, 0
	up, right, bottom, left := false, false, false, false
	_ = bottom

	helper := func() bool {
		for y != len(grid)-1 || x != len(grid[0])-1 {
			// pp.Println("x:", x, "y:", y)
			switch grid[y][x] {
			case 1:
				if left {
					x++
					if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[0]) {
						return false
					}
					if grid[y][x] != 1 && grid[y][x] != 3 && grid[y][x] != 5 {
						return false
					}
				} else {
					x--
					if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[0]) {
						return false
					}
					if grid[y][x] != 1 && grid[y][x] != 4 && grid[y][x] != 6 {
						return false
					}
				}

			case 2:
				if up {
					y++
					if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[0]) {
						return false
					}
					if grid[y][x] != 2 && grid[y][x] != 5 && grid[y][x] != 6 {
						return false
					}
				} else {
					y--
					if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[0]) {
						return false
					}
					if grid[y][x] != 2 && grid[y][x] != 3 && grid[y][x] != 4 {
						return false
					}
				}

			case 3:
				if left {
					left, up = false, true
					y++
					if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[0]) {
						return false
					}
					if grid[y][x] != 2 && grid[y][x] != 5 && grid[y][x] != 6 {
						return false
					}
				} else {
					bottom, right = false, true
					x--
					if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[0]) {
						return false
					}
					if grid[y][x] != 1 && grid[y][x] != 4 && grid[y][x] != 6 {
						return false
					}
				}

			case 4:
				if right {
					right, up = false, true
					y++
					if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[0]) {
						return false
					}
					if grid[y][x] != 2 && grid[y][x] != 5 && grid[y][x] != 6 {
						return false
					}
				} else {
					bottom, left = false, true
					x++
					if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[0]) {
						return false
					}
					if grid[y][x] != 1 && grid[y][x] != 3 && grid[y][x] != 5 {
						return false
					}
				}

			case 5:
				if up {
					up, right = false, true
					x--
					if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[0]) {
						return false
					}
					if grid[y][x] != 1 && grid[y][x] != 4 && grid[y][x] != 6 {
						return false
					}
				} else {
					left, bottom = false, true
					y--
					if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[0]) {
						return false
					}
					if grid[y][x] != 2 && grid[y][x] != 3 && grid[y][x] != 4 {
						return false
					}
				}

			case 6:
				if up {
					up, left = false, true
					x++
					if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[0]) {
						return false
					}
					if grid[y][x] != 1 && grid[y][x] != 3 && grid[y][x] != 5 {
						return false
					}
				} else {
					right, bottom = false, true
					y--
					if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[0]) {
						return false
					}
					if grid[y][x] != 2 && grid[y][x] != 3 && grid[y][x] != 4 {
						return false
					}
				}

			}
		}
		return true
	}

	switch grid[y][x] {
	case 1:
		left = true
	case 2:
		up = true
	case 3:
		left = true
	case 4:
		bottom = true
		if !helper() {
			// pp.Println("again")
			up, right, bottom, left = false, false, false, false
			right = true
			y, x = 0, 0
			return helper()
		}
	case 5:
		return false
	case 6:
		up = true
	}

	return helper()
}
