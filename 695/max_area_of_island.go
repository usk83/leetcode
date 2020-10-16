// 695. Max Area of Island
// https://leetcode.com/problems/max-area-of-island/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxAreaOfIsland([][]int{{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 1, 1}, {0, 0, 0, 1, 1}}))
}

func maxAreaOfIsland(grid [][]int) int {
	verticalLength := len(grid)
	if verticalLength == 0 {
		return 0
	}

	horizontalLength := len(grid[0])
	if horizontalLength == 0 {
		return 0
	}

	var restOfIsland int
	for i := 0; i < verticalLength; i++ {
		for j := 0; j < horizontalLength; j++ {
			restOfIsland += grid[i][j] // 0と1のみを想定した暗黙的処理
		}
	}
	if restOfIsland == 0 {
		return 0
	}

	var explore func(count, i, j int) int
	explore = func(count, i, j int) int {
		if grid[i][j] == 0 {
			return count
		}
		grid[i][j] = 0
		count++

		if i != 0 { // 上の探索
			count = explore(count, i-1, j)
		}
		if j != horizontalLength-1 { // 右の探索
			count = explore(count, i, j+1)
		}
		if i != verticalLength-1 { // 下の探索
			count = explore(count, i+1, j)
		}
		if j != 0 { // 左の探索
			count = explore(count, i, j-1)
		}

		return count
	}

	var maxAreaSize int
	for i := 0; i < verticalLength; i++ {
		for j := 0; j < horizontalLength; j++ {
			if size := explore(0, i, j); size != 0 {
				restOfIsland -= size
				if maxAreaSize < size {
					maxAreaSize = size
				}
				if restOfIsland <= maxAreaSize {
					return maxAreaSize // 帯域脱出できない
				}
			}
		}
	}

	return 0 // 本当だったらここはエラーを返す
}

func maxAreaOfIsland02(grid [][]int) int {
	verticalLength := len(grid)
	if verticalLength == 0 {
		return 0
	}

	horizontalLength := len(grid[0])
	if horizontalLength == 0 {
		return 0
	}

	checked := make([][]bool, verticalLength)
	for i := range checked {
		checked[i] = make([]bool, horizontalLength)
	}

	var walk func(count, i, j int) int
	walk = func(count, i, j int) int {
		if grid[i][j] == 0 || checked[i][j] {
			return count
		}
		checked[i][j] = true
		count++

		if i != 0 { // 上の探索
			count = walk(count, i-1, j)
		}
		if j != horizontalLength-1 { // 右の探索
			count = walk(count, i, j+1)
		}
		if i != verticalLength-1 { // 下の探索
			count = walk(count, i+1, j)
		}
		if j != 0 { // 左の探索
			count = walk(count, i, j-1)
		}

		return count
	}

	var maxAreaSize int
	for i := 0; i < verticalLength; i++ {
		for j := 0; j < horizontalLength; j++ {
			if size := walk(0, i, j); size != 0 {
				if maxAreaSize < size {
					maxAreaSize = size
				}
			}
		}
	}

	return maxAreaSize
}

func maxAreaOfIsland01(grid [][]int) int {
	verticalLength := len(grid)
	if verticalLength == 0 {
		return 0
	}

	horizontalLength := len(grid[0])
	if horizontalLength == 0 {
		return 0
	}

	checked := make([][]bool, verticalLength)
	for i := range checked {
		checked[i] = make([]bool, horizontalLength)
	}

	var restOfIsland int
	for i := 0; i < verticalLength; i++ {
		for j := 0; j < horizontalLength; j++ {
			restOfIsland += grid[i][j] // 0と1のみを想定した暗黙的処理
		}
	}
	if restOfIsland == 0 {
		return 0
	}

	var explore func(count, i, j int) int
	explore = func(count, i, j int) int {
		if grid[i][j] == 0 || checked[i][j] {
			return count
		}
		checked[i][j] = true
		count++

		if i != 0 { // 上の探索
			count = explore(count, i-1, j)
		}
		if j != horizontalLength-1 { // 右の探索
			count = explore(count, i, j+1)
		}
		if i != verticalLength-1 { // 下の探索
			count = explore(count, i+1, j)
		}
		if j != 0 { // 左の探索
			count = explore(count, i, j-1)
		}

		return count
	}

	var maxAreaSize int
	for i := 0; i < verticalLength; i++ {
		for j := 0; j < horizontalLength; j++ {
			if size := explore(0, i, j); size != 0 {
				restOfIsland -= size
				if maxAreaSize < size {
					maxAreaSize = size
				}
				if restOfIsland <= maxAreaSize {
					return maxAreaSize // 帯域脱出できない
				}
			}
		}
	}

	return 0 // 本当だったらここはエラーを返す
}
