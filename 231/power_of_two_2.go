// 231. Power of Two
// https://leetcode.com/problems/power-of-two/
package main

import (
	"math/bits"

	"github.com/k0kubun/pp"
)

// June LeetCoding Challenge Day 8
// https://leetcode.com/explore/challenge/card/june-leetcoding-challenge/540/week-2-june-8th-june-14th/3354/
func main() {
	pp.Println("=========================================")
	pp.Println(isPowerOfTwo(1))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(isPowerOfTwo(16))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(isPowerOfTwo(218))
	pp.Println(false)
	pp.Println("=========================================")
}

func isPowerOfTwo(n int) bool {
	return bits.OnesCount(uint(n)) == 1
}
