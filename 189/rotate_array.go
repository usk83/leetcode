// 189. Rotate Array
// https://leetcode.com/problems/rotate-array/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	copy(nums[2:], nums[:8])
	fmt.Println(nums)

	// pp.Println("=========================================")
	// nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// rotate(nums, 2)
	// fmt.Println(nums)
	// fmt.Println([]int{9, 10, 1, 2, 3, 4, 5, 6, 7, 8})
	// pp.Println("=========================================")
	// nums = []int{-1, -100, 3, 99}
	// rotate(nums, 2)
	// fmt.Println(nums)
	// fmt.Println([]int{3, 99, -1, -100})
	// pp.Println("=========================================")
	// nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	// rotate(nums, 3)
	// fmt.Println(nums)
	// fmt.Println([]int{10, 11, 12, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	// pp.Println("=========================================")
	// nums = []int{1, 2, 3, 4, 5, 6}
	// rotate(nums, 4)
	// fmt.Println(nums)
	// fmt.Println([]int{3, 4, 5, 6, 1, 2})
	// pp.Println("=========================================")
	// nums = []int{-1, -100, 3, 99}
	// rotate(nums, 3)
	// fmt.Println(nums)
	// fmt.Println([]int{-100, 3, 99, -1})
	// pp.Println("=========================================")
	// nums = []int{1, 2, 3, 4, 5, 6}
	// rotate(nums, 1)
	// fmt.Println(nums)
	// fmt.Println([]int{6, 1, 2, 3, 4, 5})
	// pp.Println("=========================================")
}

// Try to come up as many solutions as you can, there are at least 3 different ways to solve this problem.
// Could you do it in-place with O(1) extra space?

// func rotateX(nums []int, k int) {
// 	numsLength := len(nums)
// 	if numsLength == 0 {
// 		return
// 	}
// 	k %= numsLength
// 	if k == 0 {
// 		return
// 	}
// 	if numsLength%k == 0 {
// 		for i := k - 1; i >= 0; i-- {
// 			for j := 1; j < numsLength/k; j++ {
// 				nums[i], nums[i+k*j] = nums[i+k*j], nums[i]
// 			}
// 		}
// 	} else {
// 		for i := range nums[1:] {
// 			nums[0], nums[k*(i+1)%numsLength] = nums[k*(i+1)%numsLength], nums[0]
// 		}
// 	}
// }

// func rotate(nums []int, k int) {

// }

// Reversing and reversing
// Runtime: 48 ms, faster than 98.99% of Go online submissions for Rotate Array.
// Memory Usage: 7.6 MB, less than 66.67% of Go online submissions for Rotate Array.
// func rotate(nums []int, k int) {
// 	for h, t := 0, len(nums)-1; h < t; h, t = h+1, t-1 {
// 		nums[h], nums[t] = nums[t], nums[h]
// 	}
// 	k %= len(nums)
// 	for h, t := 0, k-1; h < t; h, t = h+1, t-1 {
// 		nums[h], nums[t] = nums[t], nums[h]
// 	}
// 	for h, t := k, len(nums)-1; h < t; h, t = h+1, t-1 {
// 		nums[h], nums[t] = nums[t], nums[h]
// 	}
// }

/*
 * 4. Partitional replacing
 * - Time Complexity: O()
 * - Space Complexity: O()
 * > Runtime: 4 ms, faster than 94.72%
 * > Memory Usage: 3.2 MB, less than 100.00%
 */
func rotate(nums []int, k int) {
	numsLength := len(nums)
	k %= numsLength
	head, tail := numsLength-k, k
	switch {
	case head == tail:
		for i := head - 1; i >= 0; i-- {
			nums[i], nums[i+head] = nums[i+head], nums[i]
		}
	case head > tail:
		pp.Println("head:", head, "tail:", tail)
		backup := make([]int, tail)
		for i := 0; i < tail; i++ {
			backup[i], nums[i] = nums[i], nums[head+i]
		}
		fmt.Println("nums:", nums)
		fmt.Println("backup:", backup)
		for i := numsLength - tail - 1; i >= tail; i-- {
			nums[i+k] = nums[i]
		}
		fmt.Println("nums:", nums)
		for i, bk := range backup {
			nums[tail+i] = bk
		}
	default:
		backup := make([]int, head)
		for i := 0; i < head; i++ {
			backup[i], nums[tail+i] = nums[tail+i], nums[i]
		}
		for i := 0; i < tail-head; i++ {
			nums[i] = nums[i+head]
		}
		for i, bk := range backup {
			nums[tail-head+i] = bk
		}
	}
}

/*
 * 3. Backing up the latter part of the array
 * - Time Complexity: O(N)
 * - Space Complexity: O(N)
 * > Runtime: 4 ms, faster than 94.72%
 * > Memory Usage: 3.3 MB, less than 100.00%
 */
// func rotate(nums []int, k int) {
// 	k %= len(nums)
// 	i := len(nums) - k
// 	latter := make([]int, k)
// 	copy(latter, nums[i:])
// 	copy(nums[k:], nums[:i])
// 	copy(nums[:k+1], latter)
// }

/*
 * 2. Cloning the whole array
 * - Time Complexity: O(N)
 * - Space Complexity: O(N)
 * > Runtime: 4 ms, faster than 94.72%
 * > Memory Usage: 3.3 MB, less than 100.00%
 */
// func rotate(nums []int, k int) {
// 	clone := make([]int, len(nums))
// 	copy(clone, nums)
// 	for i := range nums {
// 		nums[(i+k)%len(nums)] = clone[i]
// 	}
// }

/*
 * 1. Brute Force
 * - Time Complexity: O(N*K)
 * - Space Complexity: O(1)
 * > Runtime: 72 ms, faster than 21.64%
 * > Memory Usage: 3.2 MB, less than 100.00%
 */
// func rotate(nums []int, k int) {
// 	numsLength := len(nums)
// 	if numsLength <= 1 {
// 		return
// 	}
// 	for count := k % numsLength; count > 0; count-- {
// 		last := nums[numsLength-1]
// 		for i := numsLength - 1; i > 0; i-- {
// 			nums[i] = nums[i-1]
// 		}
// 		nums[0] = last
// 	}
// }
