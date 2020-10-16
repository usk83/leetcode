// 886. Possible Bipartition
// https://leetcode.com/problems/possible-bipartition/
package main

import (
	"github.com/k0kubun/pp"
)

// May LeetCoding Challenge Day 27
// https://leetcode.com/explore/challenge/card/may-leetcoding-challenge/537/week-4-may-22nd-may-28th/3342/
func main() {
	pp.Println("=========================================")
	pp.Println(possibleBipartition(4, [][]int{{1, 2}, {1, 3}, {2, 4}}))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(possibleBipartition(3, [][]int{{1, 2}, {1, 3}, {2, 3}}))
	pp.Println(false)
	pp.Println("=========================================")
	pp.Println(possibleBipartition(5, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {1, 5}}))
	pp.Println(false)
	pp.Println("=========================================")
}

// 1 <= N <= 2000
// 0 <= dislikes.length <= 10000
// 1 <= dislikes[i][j] <= N
// dislikes[i][0] < dislikes[i][1]
// There does not exist i != j for which dislikes[i] == dislikes[j].
func possibleBipartition(n int, dislikes [][]int) bool {
	graph := make([][]int, n)
	for _, dislike := range dislikes {
		graph[dislike[0]-1] = append(graph[dislike[0]-1], dislike[1]-1)
		graph[dislike[1]-1] = append(graph[dislike[1]-1], dislike[0]-1)
	}
	return true
}

func possibleBipartition(N int, dislikes [][]int) bool {
	graph := make([][]int, N)
	for _, dislike := range dislikes {
		graph[dislike[0]-1] = append(graph[dislike[0]-1], dislike[1]-1)
		graph[dislike[1]-1] = append(graph[dislike[1]-1], dislike[0]-1)
	}

	people := make([]int, N)

	var dfs func(person, neighborGroup int) bool
	dfs = func(person, neighborGroup int) bool {
		var group int
		if neighborGroup == 1 {
			group = 2
		} else {
			group = 1
		}

		// 同じだったら
		// 	return true
		// 違ったら
		// 	return false
		if people[person] != 0 {
			return people[person] == group
		}

		// 無色だったら
		// 	塗って
		// 	隣接へ進む
		people[person] = group
		for _, neighbor := range graph[person] {
			if !dfs(neighbor, group) {
				return false
			}
		}
		return true
	}

	for i, group := range people {
		if group == 0 && !dfs(i, 2) {
			return false
		}
	}

	return true
}
