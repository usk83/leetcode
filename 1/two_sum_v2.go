// 1. Two Sum
// https://leetcode.com/problems/two-sum/
package main

import "fmt"

func main() {
	fmt.Println("=========================================")
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println([]int{0, 1})
	fmt.Println("=========================================")
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println([]int{1, 2})
	fmt.Println("=========================================")
}

func twoSum(nums []int, target int) []int {
	cache := make(map[int]int, len(nums))
	for i, num := range nums {
		if index, found := cache[num]; found {
			return []int{index, i}
		}
		cache[target-num] = i
	}
	return nil
}

// func twoSum(nums []int, target int) []int {
// 	cache := make(map[int]int, len(nums))
// 	for i, num := range nums {
// 		if index, found := cache[target-num]; found {
// 			return []int{index, i}
// 		}
// 		cache[num] = i
// 	}
// 	return nil
// }

// func twoSum(nums []int, target int) []int {
// 	for i := len(nums) - 1; i >= 0; i-- {
// 		for j := i - 1; j >= 0; j-- {
// 			if nums[i]+nums[j] == target {
// 				return []int{j, i}
// 			}
// 		}
// 	}
// 	return nil
// }
