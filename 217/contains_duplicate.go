// 217. Contains Duplicate
// https://leetcode.com/problems/contains-duplicate/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(containsDuplicate([]int{1, 2, 3, 1}))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(containsDuplicate([]int{1, 2, 3, 4}))
	pp.Println(false)
	pp.Println("=========================================")
	pp.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
	pp.Println(true)
	pp.Println("=========================================")
}

func containsDuplicate(nums []int) bool {
	set := map[int]struct{}{}
	for _, num := range nums {
		if _, ok := set[num]; ok {
			return true
		}
		set[num] = struct{}{}
	}
	return false
}

// func containsDuplicate(nums []int) bool {
// 	set := map[int]bool{}
// 	for _, num := range nums {
// 		if set[num] {
// 			return true
// 		}
// 		set[num] = true
// 	}
// 	return false
// }
