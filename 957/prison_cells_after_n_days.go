// 957. Prison Cells After N Days
// https://leetcode.com/problems/prison-cells-after-n-days/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

// July LeetCoding Challenge Day 3
// https://leetcode.com/explore/featured/card/july-leetcoding-challenge/544/week-1-july-1st-july-7th/3379/
func main() {
	pp.Println("=========================================")
	fmt.Println(prisonAfterNDays([]int{0, 1, 0, 1, 1, 0, 0, 1}, 7))
	fmt.Println([]int{0, 0, 1, 1, 0, 0, 0, 0})
	pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
}

func prisonAfterNDays(cells []int, N int) []int {
	var state int
	for i := range cells {
		state |= cells[i] & 1 << uint(i+1)
	}
	fmt.Printf("%b\n", state)
	for i := 0; i < N; i++ {
		var nextState int
		for j := 2; j <= 7; j++ {
			nextState |= ((state>>uint(j-1)&1 ^ state>>uint(j+1)&1) ^ 1) << uint(j)
		}
		state = nextState
		fmt.Printf("%b\n", state)
	}
	for i := range cells {
		cells[i] = state >> uint(i+1) & 1
	}
	return cells
}
