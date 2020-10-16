// 704. Binary Search
// https://leetcode.com/problems/binary-search/
package main

import (
	"sort"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
	pp.Println(4)
	pp.Println("=========================================")
	pp.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2))
	pp.Println(-1)
	pp.Println("=========================================")
	pp.Println(search([]int{1, 2, 3, 4, 5}, 6))
	pp.Println(-1)
	pp.Println("=========================================")
}

/*
 * v2. Using `sort` package
 */
func search(nums []int, target int) int {
	if i := sort.SearchInts(nums, target); i < len(nums) && nums[i] == target {
		return i
	}
	return -1
}

/*
 * v1. Self implementation
 */
func search(nums []int, target int) int {
	start, end := 0, len(nums)
	for start < end {
		mid := (start + end) >> 1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return -1
}
