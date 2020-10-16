// 5354. Time Needed to Inform All Employees
// https://leetcode.com/contest/weekly-contest-179/problems/time-needed-to-inform-all-employees/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(numOfMinutes(1, 0, []int{-1}, []int{0}))
	pp.Println(0)
	pp.Println("=========================================")
	pp.Println(numOfMinutes(6, 2, []int{2, 2, -1, 2, 2, 2}, []int{0, 0, 1, 0, 0, 0}))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(numOfMinutes(7, 6, []int{1, 2, 3, 4, 5, 6, -1}, []int{0, 6, 5, 4, 3, 2, 1}))
	pp.Println(21)
	pp.Println("=========================================")
	pp.Println(numOfMinutes(15, 0, []int{-1, 0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6}, []int{1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}))
	pp.Println(3)
	pp.Println("=========================================")
	pp.Println(numOfMinutes(4, 2, []int{3, 3, -1, 2}, []int{0, 0, 162, 914}))
	pp.Println(1076)
	pp.Println("=========================================")
}

func numOfMinutes(n int, headID int, managers []int, informTime []int) int {
	subordinates := make([][]int, len(managers))
	for i := range subordinates {
		subordinates[i] = []int{}
	}
	for i, manager := range managers {
		if manager == -1 {
			continue
		}
		subordinates[manager] = append(subordinates[manager], i)
	}

	longest := 0

	var dfs func(id, time int)
	dfs = func(id, time int) {
		if len(subordinates[id]) == 0 {
			if time > longest {
				longest = time
			}
			return
		}

		time += informTime[id]
		for _, subordinate := range subordinates[id] {
			dfs(subordinate, time)
		}
	}

	dfs(headID, 0)

	return longest
}
