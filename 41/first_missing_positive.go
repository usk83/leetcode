// 41. First Missing Positive
// https://leetcode.com/problems/first-missing-positive/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(firstMissingPositive([]int{1, 2, 0}))
	pp.Println(3)
	pp.Println("=========================================")
	pp.Println(firstMissingPositive([]int{3, 4, -1, 1}))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(firstMissingPositive([]int{7, 8, 9, 11, 12}))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(firstMissingPositive([]int{}))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(firstMissingPositive([]int{1}))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(firstMissingPositive([]int{1, 1}))
	pp.Println(2)
	pp.Println("=========================================")
}

/*
 * v2. the clearest solution
 * ref: https://leetcode.com/discuss/topic/17071
 * - Time Complexity: O(N) (N + N)
 * - Space Complexity: O(1)
 */
func firstMissingPositive(nums []int) int {
	for i := range nums {
		for nums[i] > 0 && nums[i] <= len(nums) && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := range nums {
		if num := i + 1; num != nums[i] {
			return num
		}
	}
	return len(nums) + 1
}

/*
 * v1. the first answer (solved by referring to Hint 1)
 * - Time Complexity: O(N) (N + 2N + N)
 * - Space Complexity: O(1)
 */
// func firstMissingPositive(nums []int) int {
// 	if len(nums) == 0 {
// 		return 1
// 	}
//
// 	var min = nums[0]
// 	for _, n := range nums[1:] {
// 		if n < min {
// 			min = n
// 		}
// 	}
//
// 	visited := min - 1 // TODO: consider overflow
// 	if visited == 0 {
// 		visited = -1
// 	}
//
// 	for i := range nums {
// 		initial := true
// 		for {
// 			num := nums[i]
// 			if num == visited {
// 				break
// 			}
//
// 			if initial {
// 				nums[i] = 0
// 			} else {
// 				nums[i] = visited
// 			}
//
// 			if num < min || num <= 0 || num > len(nums) {
// 				break
// 			}
// 			i = num - 1
// 			initial = false
// 		}
// 	}
//
// 	for i, num := range nums {
// 		if num != visited {
// 			return i + 1
// 		}
// 	}
//
// 	return len(nums) + 1
// }
