// 338. Counting Bits
// https://leetcode.com/problems/counting-bits/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

// May LeetCoding Challenge Day 28
// https://leetcode.com/explore/challenge/card/may-leetcoding-challenge/537/week-4-may-22nd-may-28th/3343/
func main() {
	pp.Println("=========================================")
	fmt.Println(countBits(2))
	fmt.Println([]int{0, 1, 1})
	pp.Println("=========================================")
	fmt.Println(countBits(5))
	fmt.Println([]int{0, 1, 1, 2, 1, 2})
	pp.Println("=========================================")
}

/*
 * 2. https://leetcode.com/problems/counting-bits/discuss/79539/Three-Line-Java-Solution
 */
// func countBits(num int) []int {
// 	bits := make([]int, num+1)
// 	for i := 1; i <= num; i++ {
// 		bits[i] = bits[i>>1] + i&1 // bits[i/2] + i%2
// 	}
// 	return bits
// }

/*
 * 1. The first solution
 */
func countBits(num int) []int {
	bits := make([]int, num+1)
	var currentIndex, lastIndex int
	for i := 1; i <= num; i++ {
		bits[i] = bits[currentIndex] + 1
		if currentIndex == lastIndex {
			currentIndex, lastIndex = 0, i
		} else {
			currentIndex++
		}
	}
	return bits
}
