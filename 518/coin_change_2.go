// 518. Coin Change 2
// https://leetcode.com/problems/coin-change-2/
package main

import (
	"github.com/k0kubun/pp"
)

// June LeetCoding Challenge Day 7
// https://leetcode.com/explore/featured/card/june-leetcoding-challenge/539/week-1-june-1st-june-7th/3353/
func main() {
	pp.Println("=========================================")
	pp.Println(change(5, []int{1, 2, 5}))
	pp.Println(4)
	pp.Println("=========================================")
	pp.Println(change(3, []int{2}))
	pp.Println(0)
	pp.Println("=========================================")
	pp.Println(change(10, []int{10}))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(change(500, []int{1, 2, 5}))
	pp.Println(12701)
	pp.Println("=========================================")
}

func change(amount int, coins []int) int {
	dp := [2][]int{
		make([]int, amount+1),
		make([]int, amount+1),
	}
	dp[0][0], dp[1][0] = 1, 1
	for i := range coins {
		for j := range dp[i&1][1:] {
			dp[i&1][j+1] = dp[i&1^1][j+1]
			if j+1 >= coins[i] {
				dp[i&1][j+1] += dp[i&1][j+1-coins[i]]
			}
		}
		// fmt.Println(dp[i&1])
	}
	return dp[len(coins)&1^1][amount]
}
