// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(countServers([][]int{{1, 0}, {0, 1}}))
	pp.Println(0)
	pp.Println("=========================================")
	pp.Println(countServers([][]int{{1, 0}, {1, 1}}))
	pp.Println(3)
	pp.Println("=========================================")
	pp.Println(countServers([][]int{{1, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}))
	pp.Println(4)
	pp.Println("=========================================")
}

// m == grid.length
// n == grid[i].length
// 1 <= m <= 250
// 1 <= n <= 250
// grid[i][j] == 0 or 1
func countServers(grid [][]int) int {
	dictX := map[int]int{}
	dictY := map[int]int{}
	for i, row := range grid {
		for j, cell := range row {
			if cell == 1 {
				dictX[j] = dictX[j] + 1
				dictY[i] = dictY[i] + 1
			}
		}
	}

	var count int
	for i, row := range grid {
		for j, cell := range row {
			if cell == 1 {
				if dictX[j] > 1 {
					count++
					continue
				}
				if dictY[i] > 1 {
					count++
				}
			}
		}
	}

	return count
}
