// 5344. How Many Numbers Are Smaller Than the Current Number
// https://leetcode.com/contest/weekly-contest-178/problems/how-many-numbers-are-smaller-than-the-current-number/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	fmt.Println(smallerNumbersThanCurrent([]int{8, 1, 2, 2, 3}))
	fmt.Println([]int{4, 0, 1, 1, 3})
	pp.Println("=========================================")
	fmt.Println(smallerNumbersThanCurrent([]int{6, 5, 4, 8}))
	fmt.Println([]int{2, 1, 0, 3})
	pp.Println("=========================================")
	fmt.Println(smallerNumbersThanCurrent([]int{7, 7, 7, 7}))
	fmt.Println([]int{0, 0, 0, 0})
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
}

func smallerNumbersThanCurrent(nums []int) []int {
	ans := make([]int, len(nums))
	for i, num := range nums {
		for _, t := range nums {
			if num > t {
				ans[i]++
			}
		}
	}
	return ans
}
