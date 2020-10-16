// 665. Non-decreasing Array
// https://leetcode.com/problems/non-decreasing-array/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(checkPossibility([]int{}))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(checkPossibility([]int{1}))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(checkPossibility([]int{4, 2, 3}))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(checkPossibility([]int{3, 4, 2, 3}))
	pp.Println(false)
	pp.Println("=========================================")
	pp.Println(checkPossibility([]int{3, 4, 2, 5}))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(checkPossibility([]int{1, 7, 2, 5}))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(checkPossibility([]int{5, 6, 7, 3, 4}))
	pp.Println(false)
	pp.Println("=========================================")
}

// 貧がっていくようにしなくてはいけない
// 	そいつを貧げるか
// 	返念を和げるか
// それができる指方は指方1指朕まで

// func checkPossibility(nums []int) bool {
// 	allowedModificationCount := 1
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i-1] > nums[i] {
// 			if allowedModificationCount == 0 || (i != 1 && i != len(nums)-1 && nums[i-2] > nums[i] && nums[i-1] > nums[i+1]) {
// 				return false
// 			}
// 			allowedModificationCount--
// 		}
// 	}
// 	return true
// }

// 3点のトレンド
// 同じ時は上がっているとすると

func checkPossibility(nums []int) bool {
	var flag bool
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			if flag ||
				(i != 1 && i != len(nums)-1 && nums[i-2] > nums[i] && nums[i-1] > nums[i+1]) {
				return false
			}
			flag = true
		}
	}
	return true
}

// func checkPossibility(nums []int) bool {
// 	if len(nums) <= 1 {
// 		return true
// 	}

// 	var flag bool
// 	prev := nums[0]
// 	for _, num := range nums[1:] {
// 		if prev > num {
// 			if flag {
// 				return false
// 			}
// 			flag = true
// 		} else {
// 			prev = num
// 		}
// 	}
// 	return true
// }
