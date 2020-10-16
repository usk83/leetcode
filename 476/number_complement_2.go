// 476. Number Complement
// https://leetcode.com/problems/number-complement/
package main

import (
	"math/bits"

	"github.com/k0kubun/pp"
)

// May LeetCoding Challenge Day 4
// https://leetcode.com/explore/challenge/card/may-leetcoding-challenge/534/week-1-may-1st-may-7th/3319/
func main() {
	pp.Println("=========================================")
	pp.Println(findComplement(5))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(findComplement(1))
	pp.Println(0)
	pp.Println("=========================================")
}

/*
 * 2. XOR instead of NOT + AND
 */
// func findComplement(num int) int {
// 	return num ^ (1<<uint(bits.Len(uint(num))) - 1)
// }

/*
 * 1. The first solution
 */
func findComplement(num int) int {
	return ^num & (1<<uint(bits.Len(uint(num))) - 1)
}
