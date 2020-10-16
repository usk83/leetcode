// 279. Perfect Squares
// https://leetcode.com/problems/perfect-squares/
package main

import (
	"math"
	"sort"

	"github.com/k0kubun/pp"
)

// June LeetCoding Challenge Day 27
// https://leetcode.com/explore/challenge/card/june-leetcoding-challenge/542/week-4-june-22nd-june-28th/3373/
func main() {
	pp.Println("=========================================")
	pp.Println(numSquares(12))
	pp.Println(3)
	pp.Println("=========================================")
	pp.Println(numSquares(13))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(numSquares(17))
	pp.Println(2)
	pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
}

func numSquares(n int) int {
	// sqrts := make([]int, int(math.Sqrt(float64(n))))
	// for i := range sqrts {
	// 	sqrts[i] = (i + 1) * (i + 1)
	// }
	sqrt := int(math.Sqrt(float64(n)))
	dp := map[int]int{
		n: sqrt - 1,
	}
	var count int
LOOP:
	for {
		count++
		nextDP := map[int]int{}
		for rem, lastIndex := range dp {
			maxIndex := sort.Search(lastIndex+1, func(i int) bool {
				return (i+1)*(i+1) > rem
			}) - 1
			if rem-(maxIndex+1)*(maxIndex+1) == 0 {
				break LOOP
			}
			minIndex := sort.Search(maxIndex, func(i int) bool {
				return (i+1)*(i+1) > (rem - (maxIndex+1)*(maxIndex+1))
			})
			for i := maxIndex; i >= minIndex; i-- {
				nextDP[rem-(i+1)*(i+1)] = max(nextDP[rem-(i+1)*(i+1)], i)
			}
		}
		dp = nextDP
	}
	return count
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
