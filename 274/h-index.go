// 274. H-Index
// https://leetcode.com/problems/h-index/
package main

import (
	"sort"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(hIndex([]int{3, 0, 6, 1, 5}))
	pp.Println(3)
	pp.Println("=========================================")
	pp.Println(hIndex([]int{}))
	pp.Println(0)
	pp.Println("=========================================")
	pp.Println(hIndex([]int{100}))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(hIndex([]int{1, 2, 50, 100}))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(hIndex([]int{1, 5, 50, 100}))
	pp.Println(3)
	pp.Println("=========================================")
	pp.Println(hIndex([]int{0}))
	pp.Println(0)
	pp.Println("=========================================")
}

/**
 * 2. Count
 * - Time Complexity: O(N)
 * - Space Complexity: O(N)
 */
// func hIndex(citations []int) int {
// 	count := make([]int, len(citations)+1)
// 	for _, cita := range citations {
// 		count[min(cita, len(citations))]++
// 	}
// 	var cusum int
// 	for i := len(citations); i > 0; i-- {
// 		cusum += count[i]
// 		if i <= cusum {
// 			return i
// 		}
// 	}
// 	return 0
// }

// func min(x, y int) int {
// 	if x < y {
// 		return x
// 	}
// 	return y
// }

/**
 * 1. Sort
 * - Time Complexity: O(NlogN)
 * - Space Complexity: O(logN)
 */
func hIndex(citations []int) int {
	sort.Ints(citations)
	return sort.Search(len(citations)+1, func(i int) bool {
		return i != 0 && (i > len(citations) || citations[len(citations)-i] < i)
	}) - 1
}
