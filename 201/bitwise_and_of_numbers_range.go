// 201. Bitwise AND of Numbers Range
// https://leetcode.com/problems/bitwise-and-of-numbers-range/
package main

import (
	"math/bits"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(rangeBitwiseAnd(5, 7))
	pp.Println(4)
	pp.Println("=========================================")
	pp.Println(rangeBitwiseAnd(0, 1))
	pp.Println(0)
	pp.Println("=========================================")
}

func rangeBitwiseAnd(m int, n int) int {
	return n &^ (1<<uint(bits.Len(uint(m^n))) - 1)
}
