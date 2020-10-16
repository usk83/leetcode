// 1. Two Sum
// https://leetcode.com/problems/two-sum/
package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println([]int{0, 1})
}

func twoSum(nums []int, target int) []int {
	indicesMap := map[int][]int{}
	for i, num := range nums {
		pair, ok := indicesMap[target-num]
		if ok {
			return []int{pair[0], i}
		}
		indicesMap[num] = append(indicesMap[num], i)
	}
	return nil
}

// func twoSum(nums []int, target int) []int {
// 	indicesMap := map[int][]int{}
// 	for i, num := range nums {
// 		indicesMap[num] = append(indicesMap[num], i)
// 	}

// 	for num, i := range indicesMap {
// 		j, ok := indicesMap[target-num]
// 		if !ok {
// 			continue
// 		}
// 		if i[0] == j[0] {
// 			if len(j) >= 2 {
// 				return []int{i[0], j[1]}
// 			}
// 		} else {
// 			return []int{i[0], j[0]}
// 		}
// 	}

// 	return nil
// }

// func twoSum(nums []int, target int) []int {
// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[j] == target-nums[i] {
// 				return []int{i, j}
// 			}
// 		}
// 	}
// 	return nil
// }
