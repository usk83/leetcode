// 5501. Minimum Number of Days to Disconnect Island
// https://leetcode.com/contest/weekly-contest-204/problems/minimum-number-of-days-to-disconnect-island/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(minDays([][]int{{0, 1, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}}))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(minDays([][]int{{1, 1}}))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(minDays([][]int{{1, 0, 1, 0}}))
	pp.Println(0)
	pp.Println("=========================================")
	pp.Println(minDays([][]int{{1, 1, 0, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 0, 1, 1}, {1, 1, 0, 1, 1}}))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(minDays([][]int{{1, 1, 0, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 0, 1, 1}, {1, 1, 1, 1, 1}}))
	pp.Println(2)
	pp.Println("=========================================")
}

// 島の数を数えて
// 1つじゃないとき
// 	return 0
// 1つのとき

// 周りを0に囲まれていると考えたほうがいい

func minDays(grid [][]int) int {
	// 島を数える

	size := 0

	var count func(y, x int)
	count = func(y, x int) {
		if y < 0 || y == len(grid) || x < 0 || x == len(grid[0]) || grid[y][x] != 1 {
			return
		}
		size++
		grid[y][x] = 2
		count(y-1, x)
		count(y+1, x)
		count(y, x-1)
		count(y, x+1)
	}

	numofislands := 0
LOOP:
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == 1 {
				numofislands++
				if numofislands == 2 {
					break LOOP
				}
				count(y, x)
			}
		}
	}

	if numofislands != 1 {
		return 0
	}

	dropy, dropx := 0, 0
	rep := 2

	var countAgain func(y, x int)
	countAgain = func(y, x int) {
		if y < 0 || y == len(grid) || x < 0 || x == len(grid[0]) || grid[y][x] == 0 || grid[y][x] == rep {
			return
		}
		if y == dropy && x == dropx {
			return
		}
		grid[y][x] = rep
		countAgain(y-1, x)
		countAgain(y+1, x)
		countAgain(y, x-1)
		countAgain(y, x+1)
	}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			dropy, dropx = y, x
			rep++

			numofislands := 0
			for yy := 0; yy < len(grid); yy++ {
				for xx := 0; xx < len(grid[0]); xx++ {
					if yy == dropy && xx == dropx {
						continue
					}

					if grid[yy][xx] != 0 && grid[yy][xx] != rep {
						numofislands++
						if numofislands == 2 {
							return 1
						}
						countAgain(yy, xx)
					}
				}
			}
		}
	}

	return 2

	// 	get := func(y, x int) int {
	// 		if y < 0 || y == len(grid) || x < 0 || x == len(grid[0]) || grid[y][x] == 0 {
	// 			return 0
	// 		}
	// 		return 1
	// 	}

	// 	one, two := 0, 0
	// 	var check func(y, x int)
	// 	check = func(y, x int) {
	// 		// pp.Println("y:", y, "x:", x)

	// 		if y < 0 || y == len(grid) || x < 0 || x == len(grid[0]) || grid[y][x] != 2 {
	// 			return
	// 		}
	// 		grid[y][x] = 3

	// 		sum := get(y-1, x) + get(y+1, x) + get(y, x-1) + get(y, x+1)
	// 		if y == 0 && x == 1 {
	// 			// pp.Println(sum)
	// 		}
	// 		if sum <= 1 {
	// 			// pp.Println("y:", y, "x:", x)
	// 			// pp.Println("up:", get(y-1, x))
	// 			// pp.Println("down:", get(y+1, x))
	// 			// pp.Println("left:", get(y, x-1))
	// 			// pp.Println("right:", get(y, x+1))
	// 			one++
	// 		} else if sum == 2 {
	// 			if get(y-1, x) == get(y+1, x) {
	// 				two++
	// 			}
	// 		}

	// 		check(y-1, x)
	// 		check(y+1, x)
	// 		check(y, x-1)
	// 		check(y, x+1)
	// 	}

	// NEXT:
	// 	for y := 0; y < len(grid); y++ {
	// 		for x := 0; x < len(grid[0]); x++ {
	// 			if grid[y][x] == 2 {
	// 				check(y, x)
	// 				break NEXT
	// 			}
	// 		}
	// 	}

	// 	// pp.Println("one:", one)
	// 	// pp.Println("two:", two)

	// 	if one == 2 && size == 2 {
	// 		return 2
	// 	}

	// 	if one != 0 || two == 1 {
	// 		return 1
	// 	}

	// 	return 2
}
