// 35. Search Insert Position
// https://leetcode.com/problems/search-insert-position/
package main

import (
	"sort"

	"github.com/k0kubun/pp"
)

// June LeetCoding Challenge Day 10
// https://leetcode.com/explore/challenge/card/june-leetcoding-challenge/540/week-2-june-8th-june-14th/3356/
func main() {
	pp.Println("=========================================")
	pp.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(searchInsert([]int{1, 3, 5, 6}, 7))
	pp.Println(4)
	pp.Println("=========================================")
	pp.Println(searchInsert([]int{1, 3, 5, 6}, 0))
	pp.Println(0)
	pp.Println("=========================================")
}

/*
 * 2. self binary search solution
 */
// func searchInsert(nums []int, target int) int {
// 	l, r := 0, len(nums)
// 	for l < r {
// 		m := (l + r) >> 1
// 		if nums[m] == target {
// 			return m
// 		}
// 		if nums[m] < target {
// 			l = m + 1
// 		} else {
// 			r = m
// 		}
// 	}
// 	return l
// }
//
/* shorten */
// func searchInsert(nums []int, target int) int {
// 	l, r := 0, len(nums)
// 	for l < r {
// 		if m := (l + r) >> 1; nums[m] < target {
// 			l = m + 1
// 		} else {
// 			r = m
// 		}
// 	}
// 	return l
// }

/*
 * 1. `sort.Search` solution
 */
func searchInsert(nums []int, target int) int {
	return sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})
}
