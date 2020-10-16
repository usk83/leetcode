// 1376. Time Needed to Inform All Employees
// https://leetcode.com/problems/time-needed-to-inform-all-employees/
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

/*
 * v2. DFS from bottom to top solution
 * > Runtime: 156 ms, faster than 100.00%
 * > Memory Usage: 10.6 MB, less than 100.00%
 */
func numOfMinutes(n int, headID int, managers []int, informTime []int) int {
	if n == 1 {
		return 0
	}
	times := make([]int, n)
	times[headID] = informTime[headID]
	var dfs func(int) int
	dfs = func(id int) int {
		if times[id] != 0 {
			return times[id]
		}
		times[id] = dfs(managers[id]) + informTime[id]
		return times[id]
	}
	var maxTime int
	for id := range managers {
		maxTime = max(maxTime, dfs(id))
	}
	return maxTime
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
 * v1-1. Build tree and DFS from the head
 * > Runtime: 212 ms, faster than 100.00%
 * > Memory Usage: 12 MB, less than 100.00%
 */
// func numOfMinutes(n int, headID int, managers []int, informTime []int) int {
// 	subordinates := make([][]int, n)
// 	for employee, manager := range managers {
// 		if employee == headID {
// 			continue
// 		}
// 		subordinates[manager] = append(subordinates[manager], employee)
// 	}
// 	var dfs func(int) int
// 	dfs = func(id int) int {
// 		if len(subordinates[id]) == 0 {
// 			return 0
// 		}
// 		var maxTime int
// 		for _, subordinate := range subordinates[id] {
// 			maxTime = max(maxTime, dfs(subordinate))
// 		}
// 		return maxTime + informTime[id]
// 	}
// 	return dfs(headID)
// }

/*
 * v1. Build tree and DFS form the head
 * > Runtime: 208 ms, faster than 100.00%
 * > Memory Usage: 11.4 MB, less than 100.00%
 */
// func numOfMinutes(n int, headID int, managers []int, informTime []int) int {
// 	subordinates := make([][]int, n)
// 	for employee, manager := range managers {
// 		if employee == headID {
// 			continue
// 		}
// 		subordinates[manager] = append(subordinates[manager], employee)
// 	}
// 	var maxTime int
// 	var dfs func(int, int)
// 	dfs = func(id, time int) {
// 		if len(subordinates[id]) == 0 {
// 			if time > maxTime {
// 				maxTime = time
// 			}
// 			return
// 		}
// 		time += informTime[id]
// 		for _, subordinate := range subordinates[id] {
// 			dfs(subordinate, time)
// 		}
// 	}
// 	dfs(headID, 0)
// 	return maxTime
// }
