// 5347. Minimum Cost to Make at Least One Valid Path in a Grid
// https://leetcode.com/contest/weekly-contest-178/problems/minimum-cost-to-make-at-least-one-valid-path-in-a-grid/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(minCost([][]int{{1, 1, 1, 1}, {2, 2, 2, 2}, {1, 1, 1, 1}, {2, 2, 2, 2}}))
	pp.Println(3)
	pp.Println("=========================================")
	pp.Println(minCost([][]int{{1, 1, 3}, {3, 2, 2}, {1, 1, 4}}))
	pp.Println(0)
	pp.Println("=========================================")
	pp.Println(minCost([][]int{{1, 2}, {4, 3}}))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(minCost([][]int{{2, 2, 2}, {2, 2, 2}}))
	pp.Println(3)
	pp.Println("=========================================")
	pp.Println(minCost([][]int{{4}}))
	pp.Println(0)
	pp.Println("=========================================")
}

func minCost(grid [][]int) int {
	height, width := len(grid), len(grid[0])
	if height == 1 && width == 1 {
		return 0
	}

	mins := make([][]int, height)
	for i := range mins {
		mins[i] = make([]int, width)
	}
	mins[0][0] = 1

	points := [][2]int{[2]int{0, 0}}

	cost := 1
	for {
		cur := [][2]int{}
		for _, point := range points {
			cur = append(cur, point)
			for {
				// pp.Println(point)
				// pp.Println(width)

				switch grid[point[0]][point[1]] {
				case 1:
					point[1]++
				case 2:
					point[1]--
				case 3:
					point[0]++
				case 4:
					point[0]--
				default:
					break
				}

				// pp.Println(point[1] >= width)
				if point[0] < 0 || point[0] >= height || point[1] < 0 || point[1] >= width {
					break
				}
				if point[0]+1 == height && point[1]+1 == width {
					return cost - 1
				}
				if mins[point[0]][point[1]] != 0 {
					break
				}
				mins[point[0]][point[1]] = cost
				cur = append(cur, point)
			}
		}

		// fmt.Println(cost-1, cur)

		cost++
		next := [][2]int{}
		for _, c := range cur {
			tmp := [][2]int{
				[2]int{c[0] + 1, c[1]},
				[2]int{c[0] - 1, c[1]},
				[2]int{c[0], c[1] + 1},
				[2]int{c[0], c[1] - 1},
			}

			for _, t := range tmp {
				if t[0] < 0 || t[0] >= height || t[1] < 0 || t[1] >= width {
					continue
				}
				if t[0]+1 == height && t[1]+1 == width {
					return cost - 1
				}
				if mins[t[0]][t[1]] != 0 {
					continue
				}
				mins[t[0]][t[1]] = cost
				next = append(next, t)
			}
		}
		points = next
	}
	return cost
}
