// 5379. Stone Game III
// https://leetcode.com/problems/xxx/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(stoneGameIII([]int{1, 2, 3, 7}))
	pp.Println("Bob")
	pp.Println("=========================================")
	pp.Println(stoneGameIII([]int{1, 2, 3, -9}))
	pp.Println("Alice")
	pp.Println("=========================================")
	pp.Println(stoneGameIII([]int{1, 2, 3, 6}))
	pp.Println("Tie")
	pp.Println("=========================================")
	pp.Println(stoneGameIII([]int{1, 2, 3, -1, -2, -3, 7}))
	pp.Println("Alice")
	pp.Println("=========================================")
	pp.Println(stoneGameIII([]int{-1, -2, -3}))
	pp.Println("Tie")
	pp.Println("=========================================")
}

func stoneGameIII(stoneValue []int) string {
	return ""
}
