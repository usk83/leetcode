// 5499. Detect Pattern of Length M Repeated K or More Times
// https://leetcode.com/contest/weekly-contest-204/problems/detect-pattern-of-length-m-repeated-k-or-more-times/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(containsPattern([]int{1, 2, 4, 4, 4, 4}, 1, 3))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(containsPattern([]int{1, 2, 1, 2, 1, 1, 1, 3}, 2, 2))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(containsPattern([]int{1, 2, 1, 2, 1, 3}, 2, 3))
	pp.Println(false)
	pp.Println("=========================================")
	pp.Println(containsPattern([]int{1, 2, 3, 1, 2}, 2, 2))
	pp.Println(false)
	pp.Println("=========================================")
	pp.Println(containsPattern([]int{2, 2, 2, 2}, 2, 3))
	pp.Println(false)
	pp.Println("=========================================")
}

func containsPattern(arr []int, m int, k int) bool {
LOOP:
	for left, right := 0, m; m <= len(arr); left, right = left+1, right+1 {
		pattern := arr[left:right]
		// pp.Println("pattern:", pattern)

		l := right
		for count := 1; count < k; count++ {
			// pp.Println("count:", count)
			for i := 0; i < m; i++ {
				if l+i >= len(arr) {
					break LOOP
				}
				// pp.Println("check:", pattern[i], arr[l+i])
				if pattern[i] != arr[l+i] {
					// pp.Println("false")
					continue LOOP
				}
			}
			l += m
		}
		return true
	}
	return false
}
