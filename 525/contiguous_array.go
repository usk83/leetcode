// 525. Contiguous Array
// https://leetcode.com/problems/contiguous-array/
package main

import (
	"github.com/k0kubun/pp"
)

// May LeetCoding Challenge Day 26
// https://leetcode.com/explore/featured/card/may-leetcoding-challenge/537/week-4-may-22nd-may-28th/3341/
func main() {
	pp.Println("=========================================")
	pp.Println(findMaxLength([]int{0, 1}))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(findMaxLength([]int{0, 1, 0}))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(findMaxLength([]int{0, 1, 0, 1}))
	pp.Println(4)
	pp.Println("=========================================")
	pp.Println(findMaxLength([]int{0, 0, 1, 0, 0, 0, 1, 1}))
	pp.Println(6)
	pp.Println("=========================================")
	pp.Println(findMaxLength([]int{1, 1, 1, 1, 1, 1, 1, 1}))
	pp.Println(0)
	pp.Println("=========================================")
	pp.Println(findMaxLength([]int{0, 1, 1, 1, 1, 1, 1, 0}))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(findMaxLength([]int{}))
	pp.Println(0)
	pp.Println("=========================================")
}

/*
 * Helpers
 */
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// main

/*
 * 3. Array solution
 *   * After checked the solution
 */
func findMaxLength(nums []int) int {
}

/*
 * 2. Optimized cumulative sum and brute force
 *   - Time Complexity: O(N^2)
 *   - Space Complexity: O(N)
 */
// func findMaxLength(nums []int) int {
// 	cusum := make([]int, len(nums)+1)
// 	for i, num := range nums {
// 		cusum[i+1] = cusum[i] + num
// 	}
// 	ones := cusum[len(cusum)-1]
// 	maxSize := min(len(nums)-ones, ones) << 1
// 	for size := maxSize << 1; size > 2; size -= 2 {
// 		for i := len(nums) - size; i >= 0; i-- {
// 			if (cusum[i+size]-cusum[i])<<1 == size {
// 				return size
// 			}
// 		}
// 	}
// 	if maxSize != 0 {
// 		return 2
// 	}
// 	return 0
// }

/*
 * 1. Cumulative sum and brute force
 *   - Time Complexity: O(N^2)
 *   - Space Complexity: O(N)
 */
// func findMaxLength(nums []int) int {
// 	counts := make([][2]int, len(nums)+1)
// 	for i, num := range nums {
// 		counts[i+1] = counts[i]
// 		counts[i+1][num]++
// 	}
// 	maxSize := min(counts[len(counts)-1][0], counts[len(counts)-1][1]) << 1
// 	for size := maxSize << 1; size > 2; size -= 2 {
// 		for i := len(nums) - size; i >= 0; i-- {
// 			if counts[i+size][0]-counts[i][0] == counts[i+size][1]-counts[i][1] {
// 				return size
// 			}
// 		}
// 	}
// 	if maxSize != 0 {
// 		return 2
// 	}
// 	return 0
// }

/*
 * 0-2. WA
 */
// func findMaxLength(nums []int) int {
// 	count, counts := 0, [2]int{}
// 	for _, num := range nums {
// 		counts[num]++
// 		if counts[0] == counts[1] {
// 			count = counts[0]
// 		}
// 	}
// 	return count << 1
// }

/*
 * 0-1. WA
 */
// func findMaxLength(nums []int) int {
// 	var num, length, cur, prev int
// 	for _, n := range nums {
// 		if n == num {
// 			cur++
// 		} else {
// 			length = max(length, min(cur, prev))
// 			cur, prev = 1, cur
// 			num ^= 1
// 		}
// 	}
// 	return max(length, min(cur, prev)) << 1
// }
