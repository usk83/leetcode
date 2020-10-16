// 198. House Robber
// https://leetcode.com/problems/house-robber/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(rob([]int{1, 2, 3, 1}))
	pp.Println(4)
	pp.Println("=========================================")
	pp.Println(rob([]int{2, 7, 9, 3, 1}))
	pp.Println(12)
	pp.Println("=========================================")
	pp.Println(rob([]int{1}))
	pp.Println(1)
	pp.Println("=========================================")
}

func rob(nums []int) int {
	dp := [2]int{}
	for _, num := range nums {
		// dp[0], dp[1] = dp[1]+num, max(dp[0], dp[1])
		dp[0], dp[1] = max(dp[0], dp[1]+num), dp[0]
	}
	// return max(dp[0], dp[1])
	return dp[0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
