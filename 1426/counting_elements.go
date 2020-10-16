// Counting Elements
// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/528/week-1/3289/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(countElements([]int{1, 2, 3}))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(countElements([]int{1, 1, 3, 3, 5, 5, 7, 7}))
	pp.Println(0)
	pp.Println("=========================================")
	pp.Println(countElements([]int{1, 3, 2, 3, 5, 0}))
	pp.Println(3)
	pp.Println("=========================================")
	pp.Println(countElements([]int{1, 1, 2, 2}))
	pp.Println(2)
	pp.Println("=========================================")
}

/*
 * 1. Count each nums using map
 * - Time Complexity: O(N)
 * - Space Complexity: O(N)
 */
func countElements(arr []int) int {
	numCount := map[int]int{}
	for _, num := range arr {
		numCount[num]++
	}
	var elemCount int
	for num, count := range numCount {
		if _, ok := numCount[num+1]; ok {
			elemCount += count
		}
	}
	return elemCount
}
