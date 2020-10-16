// 169. Majority Element
// https://leetcode.com/problems/majority-element/
package main

import (
	"github.com/k0kubun/pp"
)

// May LeetCoding Challenge Day 6
// https://leetcode.com/explore/challenge/card/may-leetcoding-challenge/534/week-1-may-1st-may-7th/3321/
func main() {
	pp.Println("=========================================")
	pp.Println(majorityElement([]int{3, 2, 3}))
	pp.Println(3)
	pp.Println("=========================================")
	pp.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
	pp.Println(2)
	pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
}

func majorityElement(nums []int) int {

}

/*
 * The first solution
 */
// func majorityElement(nums []int) int {
// 	half, count := len(nums)>>1, map[int]int{}
// 	for _, num := range nums {
// 		if count[num]++; count[num] > half {
// 			return num
// 		}
// 	}
// 	return -1
// }
