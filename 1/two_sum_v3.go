// 1. Two Sum
// https://leetcode.com/problems/two-sum/
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("=========================================")
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println([]int{0, 1})
	fmt.Println("=========================================")
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println([]int{1, 2})
	fmt.Println("=========================================")
}

// suguruさんの解答を参考にしたバージョン

type elem struct {
	num   int
	index int
}

func twoSum(nums []int, target int) []int {
	arr := make([]elem, len(nums))
	for index, num := range nums {
		arr[index] = elem{num, index}
	}
	sort.Slice(arr, func(i, j int) bool { return arr[i].num < arr[j].num })

	left, right := 0, len(nums)-1
	for left < right {
		if sum := arr[left].num + arr[right].num; sum == target {
			return []int{arr[left].index, arr[right].index}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return nil
}
