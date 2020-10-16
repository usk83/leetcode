// 5440. XOR Operation in an Array
// https://leetcode.com/contest/weekly-contest-194/problems/xor-operation-in-an-array/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
}

func xorOperation(n int, start int) int {
	nums := make([]int, n)
	for i := range nums {
		nums[i] = start + 2*i
	}
	var sum int
	for _, num := range nums {
		sum ^= num
	}
	return sum
}
