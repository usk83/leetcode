// 746. Min Cost Climbing Stairs
// https://leetcode.com/problems/min-cost-climbing-stairs/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(minCostClimbingStairs([]int{10, 15}))
	pp.Println(10)
	pp.Println("=========================================")
	pp.Println(minCostClimbingStairs([]int{10, 15, 20}))
	pp.Println(15)
	pp.Println("=========================================")
	pp.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
	pp.Println(6)
	pp.Println("=========================================")
}

/*
 * cost will have a length in the range [2, 1000].
 * Every cost[i] will be an integer in the range [0, 999].
 */
func minCostClimbingStairs(cost []int) int {
	for i := range cost[2:] {
		cost[i+2] += min(cost[i], cost[i+1])
	}
	return min(cost[len(cost)-2], cost[len(cost)-1])
}

// func minCostClimbingStairs(stairs []int) int {
// 	dp := [3]int{0, stairs[0], stairs[1]}
// 	for _, stair := range stairs[2:] {
// 		dp[0], dp[1], dp[2] = dp[1], dp[2], min(dp[1], dp[2])+stair
// 	}
// 	return min(dp[1], dp[2])
// }

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
