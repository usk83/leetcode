// x. xxx
// https://leetcode.com/problems/xxx/
package main

import "github.com/k0kubun/pp"

func main() {
	pp.Println("=========================================")
	pp.Println(numWays(3, 2))
	pp.Println(4)
	pp.Println("=========================================")
	pp.Println(numWays(2, 4))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(numWays(4, 2))
	pp.Println(8)
	pp.Println("=========================================")
	pp.Println(numWays(6, 13))
	pp.Println(51)
	pp.Println("=========================================")
	pp.Println(numWays(27, 7))
	// pp.Println()
	pp.Println("=========================================")
}

// 1 <= steps <= 500
// 1 <= arrLen <= 10^6
func numWays(steps int, arrLen int) int {
	width := steps / 2
	if width >= arrLen {
		width = arrLen - 1
	}

	// 横500 縦250
	dp := make([][]int, 250)
	for i := 0; i < 250; i++ {
		dp[i] = make([]int, 500)
	}

}
