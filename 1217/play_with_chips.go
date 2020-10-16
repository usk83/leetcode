// 1217. Play with Chips
// https://leetcode.com/problems/play-with-chips/
package main

import "github.com/k0kubun/pp"

func main() {
	pp.Println(minCostToMoveChips([]int{1, 2, 3}))
	pp.Println(1)

	pp.Println(minCostToMoveChips([]int{2, 2, 2, 3, 3}))
	pp.Println(2)
}

func minCostToMoveChips(chips []int) int {
	var evens, odds int
	for _, c := range chips {
		if c&1 == 0 {
			evens++
		} else {
			odds++
		}
	}

	if evens < odds {
		return evens
	}
	return odds
}
