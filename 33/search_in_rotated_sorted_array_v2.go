// 33. Search in Rotated Sorted Array
// https://leetcode.com/problems/search-in-rotated-sorted-array/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	pp.Println(4)
	pp.Println("=========================================")
	pp.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	pp.Println(-1)
	pp.Println("=========================================")
	pp.Println(search([]int{}, 5))
	pp.Println(-1)
	pp.Println("=========================================")
	pp.Println(search([]int{1}, 0))
	pp.Println(-1)
	pp.Println("=========================================")
	pp.Println(search([]int{16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, 1))
	pp.Println(13)
	pp.Println("=========================================")
	pp.Println(search([]int{1, 3}, 0))
	pp.Println(-1)
	pp.Println("=========================================")
	pp.Println(search([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28}, 26))
	pp.Println(25)
	pp.Println("=========================================")
	pp.Println(search([]int{3, 5, 1}, 5))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(search([]int{8, 9, 2, 3, 4}, 9))
	pp.Println(1)
	pp.Println("=========================================")
}

// Returns -1 if target is on the left side
// Returns 0 if target is found
// Returns 1 if  target is on the right side
type compFunc func(i int) int

func binarySearch(start, end int, comp compFunc) int {
	for start <= end {
		middle := start + (end-start)/2
		if result := comp(middle); result == 0 {
			return middle
		} else if result > 0 {
			start = middle + 1
		} else {
			end = middle - 1
		}
	}
	return -1
}

// pivotを探す
// 折れ曲がる前の最後の要素をpivotとする
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	lastIndex := len(nums) - 1
	pivot := binarySearch(0, lastIndex, func(i int) int {
		if i != lastIndex && nums[i] > nums[i+1] {
			return 0
		}
		if nums[i] >= nums[0] {
			return 1
		}
		return -1
	})

	var start, end int
	if lastValue := nums[lastIndex]; lastValue < target {
		start = 0
		end = pivot
	} else {
		start = pivot + 1
		end = lastIndex
	}

	return binarySearch(start, end, func(i int) int {
		if nums[i] == target {
			return 0
		}
		if target > nums[i] {
			return 1
		}
		return -1
	})
}
