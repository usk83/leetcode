// 137. Single Number II
// https://leetcode.com/problems/single-number-ii/
package main

import (
	"github.com/k0kubun/pp"
)

// June LeetCoding Challenge Day 22
// https://leetcode.com/explore/challenge/card/june-leetcoding-challenge/542/week-4-june-22nd-june-28th/3368/
func main() {
	pp.Println("=========================================")
	pp.Println(singleNumber([]int{2, 2, 3, 2}))
	pp.Println(3)
	pp.Println("=========================================")
	pp.Println(singleNumber([]int{0, 1, 0, 1, 0, 1, 99}))
	pp.Println(99)
	pp.Println("=========================================")
	pp.Println(singleNumber([]int{1}))
	pp.Println(1)
	pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
}

func singleNumber(nums []int) int {
	var once, twice int
	for _, num := range nums {
		once = once ^ num & ^twice
		twice = twice ^ num & ^once
	}
	return once
}

func singleNumber(nums []int) int {
	var twice, once int
	for _, num := range nums {
		twice, once = twice&^once&^num|^twice&once&num, once^num&^twice
	}
	return once
}

/**
 * ref: https://leetcode.com/problems/single-number-ii/discuss/43296/An-General-Way-to-Handle-All-this-sort-of-questions.
 */
// func singleNumber(nums []int) int {
// 	var twice, once int
// 	for _, num := range nums {
// 		twice, once = twice&^once&^num|^twice&once&num, ^twice&once&^num|^twice&^once&num
// 	}
// 	return once
// }
