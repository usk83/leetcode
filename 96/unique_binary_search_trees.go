// 96. Unique Binary Search Trees
// https://leetcode.com/problems/unique-binary-search-trees/
package main

import (
	"github.com/k0kubun/pp"
)

// June LeetCoding Challenge Day 24
// https://leetcode.com/explore/challenge/card/june-leetcoding-challenge/542/week-4-june-22nd-june-28th/3370/
func main() {
	expected := []int{
		1, 2, 5, 14, 42, 132, 429,
	}
	pp.Println("=========================================")
	for i, want := range expected {
		pp.Println(numTrees(i + 1))
		pp.Println(want)
		pp.Println("=========================================")
	}
}

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		div, mod := i>>1, i&1
		for j := 0; j < div; j++ {
			dp[i] += dp[j] * dp[i-1-j]
		}
		dp[i] <<= 1
		if mod == 1 {
			dp[i] += dp[div] * dp[div]
		}
	}
	return dp[n]
}
