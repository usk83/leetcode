// 5364. Create Target Array in the Given Order
// https://leetcode.com/contest/weekly-contest-181/problems/create-target-array-in-the-given-order/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	fmt.Println(createTargetArray([]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1}))
	fmt.Println([]int{0, 4, 1, 3, 2})
	pp.Println("=========================================")
	fmt.Println(createTargetArray([]int{1, 2, 3, 4, 0}, []int{0, 1, 2, 3, 0}))
	fmt.Println([]int{0, 1, 2, 3, 4})
	pp.Println("=========================================")
	fmt.Println(createTargetArray([]int{1}, []int{0}))
	fmt.Println([]int{1})
	pp.Println("=========================================")
}

func createTargetArray(nums []int, index []int) []int {
	result := make([]int, len(nums))
	for key, i := range index {
		new := make([]int, 0, len(nums))
		for j := 0; j < i; j++ {
			new = append(new, result[j])
		}
		new = append(new, nums[key])
		for j := i + 1; j < len(nums); j++ {
			new = append(new, result[j-1])
		}
		result = new
	}
	return result
}
