// 1051. Height Checker
// https://leetcode.com/problems/height-checker/
package main

import (
	"sort"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(heightChecker([]int{1, 1, 4, 2, 1, 3}))
	pp.Println(3)
	pp.Println("=========================================")
	pp.Println(heightChecker([]int{2, 1, 1, 1}))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(heightChecker([]int{4, 2, 1, 1, 1}))
	pp.Println(4)
	pp.Println("=========================================")
	pp.Println(heightChecker([]int{1, 1, 4, 2, 3}))
	pp.Println(3)
	pp.Println("=========================================")
}

func heightChecker(heights []int) int {
	clone := make([]int, len(heights))
	copy(clone, heights)

	sort.Ints(clone)

	var count int
	for i := range heights {
		if heights[i] != clone[i] {
			count++
		}
	}

	return count
}
