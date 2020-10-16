// 62. Unique Paths
// https://leetcode.com/problems/unique-paths/
package main

import (
	"github.com/k0kubun/pp"
)

// June LeetCoding Challenge Day 29
// https://leetcode.com/explore/challenge/card/june-leetcoding-challenge/543/week-5-june-29th-june-30th/3375/
func main() {
	pp.Println("=========================================")
	pp.Println(uniquePaths(3, 2))
	pp.Println(3)
	pp.Println("=========================================")
	pp.Println(uniquePaths(7, 3))
	pp.Println(28)
	pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
}

func uniquePaths(m int, n int) int {
	if m > n {
		m, n = n, m
	}
	dp := [2][]int{
		make([]int, m),
		make([]int, m),
	}
	for i := range dp[0] {
		dp[0][i] = 1
	}
	dp[1][0] = 1
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i&1][j] = dp[i&1^1][j] + dp[i&1][j-1]
		}
	}
	return dp[n&1^1][m-1]
}
