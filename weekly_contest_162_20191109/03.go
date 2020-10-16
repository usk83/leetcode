package main

import "github.com/k0kubun/pp"

func main() {
	pp.Println("=========================================")
	pp.Println(closedIsland([][]int{{1, 1, 1, 1, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 1, 0}, {1, 0, 1, 0, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 0, 1}, {1, 1, 1, 1, 1, 1, 1, 0}}))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(closedIsland([][]int{{0, 0, 1, 0, 0}, {0, 1, 0, 1, 0}, {0, 1, 1, 1, 0}}))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(closedIsland([][]int{{1, 1, 1, 1, 1, 1, 1}, {1, 0, 0, 0, 0, 0, 1}, {1, 0, 1, 1, 1, 0, 1}, {1, 0, 1, 0, 1, 0, 1}, {1, 0, 1, 1, 1, 0, 1}, {1, 0, 0, 0, 0, 0, 1}, {1, 1, 1, 1, 1, 1, 1}}))
	pp.Println(2)
	pp.Println("=========================================")
}

// 1 <= grid.length, grid[0].length <= 100
// 0 <= grid[i][j] <=1
func closedIsland(grid [][]int) int {
	// 陸地をたどって
	// 海にしていく
	// はしが含まれていた場合はno count
	// カウントを返却する

	yLen := len(grid)
	if yLen == 0 {
		return 0
	}
	xLen := len(grid[0])
	if xLen == 0 {
		return 0
	}

	var explore func(x, y int) bool
	explore = func(x, y int) bool {
		grid[y][x] = 1
		flag := true
		if x == 0 || x == xLen-1 || y == 0 || y == yLen-1 {
			flag = false
		}

		// 上
		if y != 0 {
			if grid[y-1][x] == 0 && !explore(x, y-1) {
				flag = false
			}
		}
		// 右
		if x != xLen-1 {
			if grid[y][x+1] == 0 && !explore(x+1, y) {
				flag = false
			}
		}
		// 下
		if y != yLen-1 {
			if grid[y+1][x] == 0 && !explore(x, y+1) {
				flag = false
			}
		}
		// 左
		if x != 0 {
			if grid[y][x-1] == 0 && !explore(x-1, y) {
				flag = false
			}
		}

		return flag
	}

	var count int
	for y := 0; y < yLen; y++ {
		for x := 0; x < xLen; x++ {
			if grid[y][x] == 1 {
				continue
			}
			if explore(x, y) {
				count++
			}
		}
	}
	return count
}
