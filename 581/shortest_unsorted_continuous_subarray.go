// 581. Shortest Unsorted Continuous Subarray
// https://leetcode.com/problems/shortest-unsorted-continuous-subarray/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))
	pp.Println(5)
	pp.Println("=========================================")
	pp.Println(findUnsortedSubarray([]int{2, 3, 10, 4, 5, 1, 6, 7, 8, 9}))
	pp.Println(10)
	pp.Println("=========================================")
	pp.Println(findUnsortedSubarray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	pp.Println(0)
	pp.Println("=========================================")
	pp.Println(findUnsortedSubarray([]int{1}))
	pp.Println(0)
	pp.Println("=========================================")
}

/*
 * v2-1. find sorted subarray from the head and the tail
 * - Time Complexity: O(N)
 * - Space Complexity: O(1)
 * > Runtime: 24 ms, faster than 96.91%
 * > Memory Usage: 6.2 MB, less than 100.00%
 */
func findUnsortedSubarray(nums []int) int {
	start, end := len(nums)-1, 0
	min, max := nums[start], nums[end]
	for i := start - 1; i >= 0; i-- {
		if nums[i] <= min {
			min = nums[i]
		} else {
			start = i
		}
	}
	if start == len(nums)-1 {
		return 0
	}
	for i := end + 1; i < len(nums); i++ {
		if nums[i] >= max {
			max = nums[i]
		} else {
			end = i
		}
	}
	return end - start + 1
}

/*
 * v2. find sorted subarray from the head and the tail
 * - Time Complexity: O(N)
 * - Space Complexity: O(1)
 * > Runtime: 24 ms, faster than 96.91%
 * > Memory Usage: 6.2 MB, less than 100.00%
 */
// func findUnsortedSubarray(nums []int) int {
// 	var head, tail int
// 	for head = 0; head < len(nums)-1 && nums[head] <= nums[head+1]; head++ {
// 	}
// 	if head == len(nums)-1 {
// 		return 0
// 	}
// 	for i := head + 1; i < len(nums); i++ {
// 		for head >= 0 && nums[head] > nums[i] {
// 			head--
// 		}
// 		if head < 0 {
// 			break
// 		}
// 	}
// 	for tail = len(nums) - 1; nums[tail] >= nums[tail-1]; tail-- {
// 	}
// 	for i := tail - 1; i >= 0; i-- {
// 		for tail < len(nums) && nums[tail] < nums[i] {
// 			tail++
// 		}
// 		if tail == len(nums) {
// 			break
// 		}
// 	}
// 	return tail - head - 1
// }

/*
 * v1. clone array, sort and compare
 * - Time Complexity: O(NlogN)
 * - Space Complexity: O(N)
 * > Runtime: 36 ms, faster than 60.82%
 * > Memory Usage: 6.2 MB, less than 100.00%
 * Time: 09:32
 */
// func findUnsortedSubarray(nums []int) int {
// 	sorted := make([]int, len(nums))
// 	copy(sorted, nums)
// 	sort.Ints(sorted)
// 	var head, tail int
// 	for head = 0; head < len(nums)-1 && nums[head] == sorted[head]; head++ {
// 	}
// 	if head == len(nums)-1 {
// 		return 0
// 	}
// 	for tail = len(nums) - 1; nums[tail] == sorted[tail]; tail-- {
// 	}
// 	return tail - head + 1
// }
