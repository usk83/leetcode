package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	fmt.Println(shiftGrid([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 1))
	fmt.Println(shiftGrid([][]int{{3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}, {12, 0, 21, 13}}, 4))
	fmt.Println(shiftGrid([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 9))
	pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
}

func shiftGrid(grid [][]int, k int) [][]int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return grid
	}

	// jump := func(x, y int) (int, int) {
	// }

	yLen := len(grid)
	xLen := len(grid[0])
	nextY := k / xLen
	nextX := k % xLen
	for nextY >= yLen {
		nextY -= yLen
	}

	result := make([][]int, yLen)
	for i := 0; i < yLen; i++ {
		result[i] = make([]int, xLen)
	}

	// prev := grid[0][0]
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			result[nextY][nextX] = grid[i][j]
			nextX++
			if nextX == xLen {
				nextX = 0
				nextY++
			}
			if nextY == yLen {
				nextY = 0
			}
		}
	}

	return result

	// for count := xLen * yLen; count >= 0; count-- {

	// 	grid[nextY][nextX], prev = prev, grid[nextY][nextX]

	// 	nextX++
	// 	if nextX == xLen {
	// 		nextX = 0
	// 		nextY++
	// 	}
	// 	if nextY == yLen {
	// 		nextY = 0
	// 	}
	// 	// for i := 0; i < len(grid); i++ {
	// 	// 	for j := 0; j < len(grid[0]); j++ {
	// 	// 		prev = grid[i][j]
	// 	// 		if i == 0 && j == 0 {
	// 	// 			continue
	// 	// 		}
	// 	// 	}
	// 	// }
	// }

	// grid[0][0] = prev

	// return grid
}
