// 207. Course Schedule
// https://leetcode.com/problems/course-schedule/
package main

import (
	"github.com/k0kubun/pp"
)

// May LeetCoding Challenge Day 29
// https://leetcode.com/explore/challenge/card/may-leetcoding-challenge/538/week-5-may-29th-may-31st/3344/
func main() {
	pp.Println("=========================================")
	pp.Println(canFinish(2, [][]int{{1, 0}}))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))
	pp.Println(false)
	pp.Println("=========================================")
	pp.Println(canFinish(1, [][]int{}))
	pp.Println(true)
	pp.Println("=========================================")
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	inDegrees := make([]int, numCourses)
	for _, prerequisite := range prerequisites {
		graph[prerequisite[0]] = append(graph[prerequisite[0]], prerequisite[1])
		inDegrees[prerequisite[1]]++
	}

	queue := make([]int, 0, numCourses)
	for i, inDegree := range inDegrees {
		if inDegree == 0 {
			queue = append(queue, i)
			numCourses--
		}
	}

	for len(queue) != 0 {
		for _, adj := range graph[queue[0]] {
			inDegrees[adj]--
			if inDegrees[adj] == 0 {
				numCourses--
				if numCourses == 0 {
					return true
				}
				queue = append(queue, adj)
			}
		}
		queue = queue[1:]
	}

	return numCourses == 0
}
