// 75. Sort Colors
// https://leetcode.com/problems/sort-colors/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

// June LeetCoding Challenge Day 11
// https://leetcode.com/explore/challenge/card/june-leetcoding-challenge/540/week-2-june-8th-june-14th/3357/
func main() {
	var nums []int
	pp.Println("=========================================")
	nums = []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
	fmt.Println([]int{0, 0, 1, 1, 2, 2})
	pp.Println("=========================================")
	nums = []int{2, 0, 1}
	sortColors(nums)
	fmt.Println(nums)
	fmt.Println([]int{0, 1, 2})
	pp.Println("=========================================")
	nums = []int{2, 1, 0, 1, 2}
	sortColors(nums)
	fmt.Println(nums)
	fmt.Println([]int{0, 1, 1, 2, 2})
	pp.Println("=========================================")
}

/* a.k.a Dutch National Flag Problem */

/**
 * 2. Three pointers solution
 * - Time Complexity: O(N)
 * - Space Complexity: O(1)
 *
 * - Runtime: 0 ms, faster than 100.00%
 * - Memory Usage: 2 MB, less than 98.96%
 */
func sortColors(nums []int) {
	for l, i, r := 0, 0, len(nums)-1; i <= r; {
		switch nums[i] {
		case 0:
			nums[i], nums[l] = nums[l], nums[i]
			l++
			i++
		case 1:
			i++
		case 2:
			nums[i], nums[r] = nums[r], nums[i]
			r--
		}
	}
}

/* shorten version */
// func sortColors(nums []int) {
// 	for l, i, r := 0, 0, len(nums)-1; i <= r; i++ {
// 		if nums[i] == 0 {
// 			nums[i], nums[l] = nums[l], nums[i]
// 			l++
// 		} else if nums[i] == 2 {
// 			nums[i], nums[r] = nums[r], nums[i]
// 			r--
// 			i--
// 		}
// 	}
// }

/**
 * 1. Bucket sort
 * - Time Complexity: O(N)
 * - Space Complexity: O(1)
 *
 * - Runtime: 0 ms, faster than 100.00%
 * - Memory Usage: 2.1 MB, less than 72.92%
 */
// func sortColors(nums []int) {
// 	var buckets [3]int
// 	for _, num := range nums {
// 		buckets[num]++
// 	}
// 	for i := range nums {
// 		if i < buckets[0] {
// 			nums[i] = 0
// 		} else if i < buckets[0]+buckets[1] {
// 			nums[i] = 1
// 		} else {
// 			nums[i] = 2
// 		}
// 	}
// }
//
// func sortColors(nums []int) {
// 	var buckets [3]int
// 	for _, num := range nums {
// 		buckets[num]++
// 	}
// 	var index int
// 	for i := range buckets {
// 		for count := buckets[i]; count > 0; count-- {
// 			nums[index] = i
// 			index++
// 		}
// 	}
// }
