// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	// pp.Println(shortestPath([][]int{
	// 	{0, 0, 0},
	// 	{0, 0, 0},
	// 	{0, 0, 0},
	// 	{0, 0, 0},
	// 	{0, 0, 0}}, 1))
	pp.Println("=========================================")
	pp.Println(shortestPath([][]int{
		{0, 0, 0},
		{1, 1, 0},
		{0, 0, 0},
		{0, 1, 1},
		{0, 0, 0}}, 1))
	pp.Println(6)
	pp.Println("=========================================")
	pp.Println(shortestPath([][]int{
		{0, 1, 1},
		{1, 1, 1},
		{1, 0, 0}}, 1))
	pp.Println(-1)
	pp.Println("=========================================")
}

// Constraints:
//     grid.length == m
//     grid[0].length == n
//     1 <= m, n <= 40
//     1 <= k <= m*n
//     grid[i][j] == 0 or 1
//     grid[0][0] == grid[m-1][n-1] == 0
func shortestPath(grid [][]int, k int) int {
	goalX, goalY := len(grid[0])-1, len(grid)-1
	minStep := -1
	visited := map[[2]int][][2]int{}          // step, delete
	positions := [][4]int{[4]int{0, 0, 0, 0}} // x, y, step, delete
	for len(positions) != 0 {
		nextPositions := [][4]int{}
		for _, position := range positions {
			x, y, step, count := position[0], position[1], position[2], position[3]
			step++
			// 左上から進めていく
			//  4方向
			// step数と消した数を記録していく
			// step数は少ないほどよい
			// 消した数も少ないほうが良い
			// 同じところにたどり着いたとき
			//  step数と消した数がともに少ないか同じ場合は終了

			// 上
			if y != 0 && (grid[y-1][x] == 0 || count < k) {
				if x == goalX && y-1 == goalY {
					minStep = min(minStep, step)
				} else {
					flag := true
					if prevs, found := visited[[2]int{x, y - 1}]; found {
						for _, prev := range prevs {
							if step >= prev[0] && count >= prev[1] {
								flag = false
								break
							}
						}
					}
					if flag {
						count, x, y, step = count+grid[y][x], x, y-1, step+1
						visited[[2]int{x, y}] = append(visited[[2]int{x, y}], [2]int{step, count})
						nextPositions = append(nextPositions, [4]int{x, y, step, count})
					}
				}
			}
			// 右
			if x != goalX && (grid[y][x+1] == 0 || count < k) {
				if x+1 == goalX && y == goalY {
					minStep = min(minStep, step)
				} else {
					flag := true
					if prevs, found := visited[[2]int{x + 1, y}]; found {
						for _, prev := range prevs {
							if step >= prev[0] && count >= prev[1] {
								flag = false
								break
							}
						}
					}
					if flag {
						count, x, y, step = count+grid[y][x], x+1, y, step+1
						visited[[2]int{x, y}] = append(visited[[2]int{x, y}], [2]int{step, count})
						nextPositions = append(nextPositions, [4]int{x, y, step, count})
					}
				}
			}
			// 下
			if y != goalY && (grid[y+1][x] == 0 || count < k) {
				if x == goalX && y+1 == goalY {
					minStep = min(minStep, step)
				} else {
					flag := true
					if prevs, found := visited[[2]int{x, y + 1}]; found {
						for _, prev := range prevs {
							if step >= prev[0] && count >= prev[1] {
								flag = false
								break
							}
						}
					}
					if flag {
						count, x, y, step = count+grid[y][x], x, y+1, step+1
						visited[[2]int{x, y}] = append(visited[[2]int{x, y}], [2]int{step, count})
						nextPositions = append(nextPositions, [4]int{x, y, step, count})
					}
				}
			}
			// 左
			if x != 0 && (grid[y][x-1] == 0 || count < k) {
				if x-1 == goalX && y == goalY {
					minStep = min(minStep, step)
				} else {
					flag := true
					if prevs, found := visited[[2]int{x - 1, y}]; found {
						for _, prev := range prevs {
							if step >= prev[0] && count >= prev[1] {
								flag = false
								break
							}
						}
					}
					if flag {
						count, x, y, step = count+grid[y][x], x-1, y, step+1
						visited[[2]int{x, y}] = append(visited[[2]int{x, y}], [2]int{step, count})
						nextPositions = append(nextPositions, [4]int{x, y, step, count})
					}
				}
			}
		}

		pp.Println("=== DEBUG START ======================================")
		pp.Println(nextPositions)
		pp.Println("=== DEBUG END ========================================")

		positions = nextPositions
	}

	return minStep
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
