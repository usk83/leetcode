// 33. Search in Rotated Sorted Array
// https://leetcode.com/problems/search-in-rotated-sorted-array/
package main

import "github.com/k0kubun/pp"

func main() {
	pp.Println("=========================================")
	pp.Println(search([]int{4, 5, 6, 7, 0, 1}, 0))
	pp.Println(4)
	pp.Println("=========================================")
}

// ref: https://leetcode.com/problems/search-in-rotated-sorted-array/discuss/14435/Clever-idea-making-it-simple
func search(nums []int, target int) int {
	start, end := 0, len(nums)
	for start < end {
		middle := start + (end-start)/2
		current := nums[middle]
		if (current < nums[0]) != (target < nums[0]) {
			if target < nums[0] {
				current = target - 1
			} else {
				current = target + 1
			}
		}
		if current == target {
			return middle
		}
		if current < target {
			start = middle + 1
		} else {
			end = middle
		}
	}
	return -1
}
