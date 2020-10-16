// 5426. Reorder Routes to Make All Paths Lead to the City Zero
// https://leetcode.com/contest/weekly-contest-191/problems/reorder-routes-to-make-all-paths-lead-to-the-city-zero/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(minReorder(6, [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}))
	pp.Println(3)
	pp.Println("=========================================")
	pp.Println(minReorder(5, [][]int{{1, 0}, {1, 2}, {3, 2}, {3, 4}}))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(minReorder(3, [][]int{{1, 0}, {2, 0}}))
	pp.Println(0)
	pp.Println("=========================================")
}

func minReorder(n int, connections [][]int) int {
	roads := map[int]map[int]bool{}
	for _, connection := range connections {
		if roads[connection[0]] == nil {
			roads[connection[0]] = map[int]bool{}
		}
		if roads[connection[1]] == nil {
			roads[connection[1]] = map[int]bool{}
		}
		roads[connection[0]][connection[1]] = false
		roads[connection[1]][connection[0]] = true
	}

	var cost int
	var dfs func(city, from int)
	dfs = func(city, from int) {
		for to, free := range roads[city] {
			if to == from {
				continue
			}
			if !free {
				cost++
			}
			dfs(to, city)
		}
	}

	for to, free := range roads[0] {
		if !free {
			cost++
		}
		dfs(to, 0)
	}

	return cost
}
