// 1219. Path with Maximum Gold
// https://leetcode.com/problems/path-with-maximum-gold/
package main

import "github.com/k0kubun/pp"

func main() {
	pp.Println(getMaximumGold([][]int{{0, 6, 0}, {5, 8, 7}, {0, 9, 0}}))
	pp.Println(24)

	pp.Println(getMaximumGold([][]int{{1, 0, 7}, {2, 0, 6}, {3, 4, 5}, {0, 3, 0}, {9, 0, 20}}))
	pp.Println(28)

	pp.Println(getMaximumGold([][]int{{5, 3, 36, 26, 27}, {11, 31, 23, 30, 4}, {5, 7, 21, 38, 10}, {39, 30, 10, 17, 13}, {16, 0, 0, 26, 1}, {25, 0, 10, 0, 0}}))
	pp.Println(454)
}

// 1 <= grid.length, grid[i].length <= 15
// 0 <= grid[i][j] <= 100
// There are at most 25 cells containing gold.
func copyMap(m map[[2]int]struct{}) map[[2]int]struct{} {
	newMap := map[[2]int]struct{}{}
	for k, v := range m {
		newMap[k] = v
	}
	return newMap
}

func getMaximumGold(grid [][]int) int {
	yLength, xLength := len(grid), len(grid[0])

	var walk func(i, j, c int, visited map[[2]int]struct{}) int
	walk = func(i, j, c int, visited map[[2]int]struct{}) int {
		if grid[i][j] == 0 {
			return c
		}

		if _, found := visited[[2]int{i, j}]; found {
			return c
		}

		visited[[2]int{i, j}] = struct{}{}
		c += grid[i][j]

		up, right, down, left := 0, 0, 0, 0

		if i != 0 { // 上の探索
			up = walk(i-1, j, c, visited)
		}
		if j != xLength-1 { // 右の探索
			right = walk(i, j+1, c, visited)
		}
		if i != yLength-1 { // 下の探索
			down = walk(i+1, j, c, visited)
		}
		if j != 0 { // 左の探索
			left = walk(i, j-1, c, visited)
		}

		delete(visited, [2]int{i, j})

		if up > c {
			c = up
		}
		if right > c {
			c = right
		}
		if down > c {
			c = down
		}
		if left > c {
			c = left
		}

		return c
	}

	var max int
	for i := range grid {
		for j := range grid[i] {
			visited := map[[2]int]struct{}{}
			count := walk(i, j, 0, visited)
			if count > max {
				max = count
			}

		}
	}

	return max
}
