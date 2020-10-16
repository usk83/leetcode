// 275. H-Index II
// https://leetcode.com/problems/h-index-ii/
package main

import (
	"sort"

	"github.com/k0kubun/pp"
)

// June LeetCoding Challenge Day 18
// https://leetcode.com/explore/featured/card/june-leetcoding-challenge/541/week-3-june-15th-june-21st/3364/
func main() {
	pp.Println("=========================================")
	pp.Println(hIndex([]int{0, 1, 3, 5, 6}))
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

func hIndex(citations []int) int {
	return sort.Search(len(citations)+1, func(i int) bool {
		return i != 0 && (i > len(citations) || citations[len(citations)-i] < i)
	}) - 1
}
