// 220. Contains Duplicate III
// https://leetcode.com/problems/contains-duplicate-iii/
package main

import (
	"github.com/k0kubun/pp"
)

// September LeetCoding Challenge Day 2
// https://leetcode.com/explore/challenge/card/september-leetcoding-challenge/554/week-1-september-1st-september-7th/3446/
func main() {
	pp.Println("=========================================")
	pp.Println(containsNearbyAlmostDuplicate([]int{1, 2, 3, 1}, 3, 0))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(containsNearbyAlmostDuplicate([]int{1, 0, 1, 1}, 1, 2))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(containsNearbyAlmostDuplicate([]int{1, 5, 9, 1, 5, 9}, 2, 3))
	pp.Println(false)
	pp.Println("=========================================")
	pp.Println(containsNearbyAlmostDuplicate([]int{-3, 3}, 2, 4))
	pp.Println(false)
	pp.Println("=========================================")
	pp.Println(containsNearbyAlmostDuplicate([]int{4, 2}, 2, 1))
	pp.Println(false)
	pp.Println("=========================================")
	pp.Println(containsNearbyAlmostDuplicate([]int{1, 2}, 0, 1))
	pp.Println(false)
	pp.Println("=========================================")
}

/*
 * Using buckets
 * ref: https://leetcode.com/problems/contains-duplicate-iii/discuss/61639/JavaPython-one-pass-solution-O(n)-time-O(n)-space-using-buckets
 * - Time Complexity: O(N)
 * - Space Complexity: O(N)
 */
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if t < 0 {
		return false
	}
	m := map[int]int{}
	width := t + 1
	for i, num := range nums {
		key := num / width
		if same, ok := m[key]; ok && max(num, same)-min(num, same) <= t {
			return true
		}
		if prev, ok := m[key-1]; ok && num-prev <= t {
			return true
		}
		if next, ok := m[key+1]; ok && next-num <= t {
			return true
		}
		if outdatedIndex := i - k; outdatedIndex >= 0 {
			delete(m, nums[outdatedIndex]/width)
		}
		m[key] = num
	}
	return false
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/*
 * Brute-force
 * - Time Complexity: O(N^2)
 * - Space Complexity: O(1)
 */
// func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
// 	for i, num1 := range nums {
// 		for _, num2 := range nums[i+1 : min(len(nums), i+1+k)] {
// 			if diff := num1 - num2; -t <= diff && diff <= t {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// func min(x, y int) int {
// 	if x < y {
// 		return x
// 	}
// 	return y
// }
