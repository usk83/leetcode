// 1389. Create Target Array in the Given Order
// https://leetcode.com/problems/create-target-array-in-the-given-order/
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
	fmt.Println(createTargetArray([]int{}, []int{}))
	fmt.Println([]int{})
	pp.Println("=========================================")
	fmt.Println(createTargetArray([]int{1, 2, 3, 4, 5}, []int{0, 1, 2, 3, 4}))
	fmt.Println([]int{1, 2, 3, 4, 5})
	pp.Println("=========================================")
	fmt.Println(createTargetArray([]int{4, 2, 4, 3, 2}, []int{0, 0, 1, 3, 1}))
	fmt.Println([]int{2, 2, 4, 4, 3})
	pp.Println("=========================================")
}

// 1 <= nums.length, index.length <= 100
// nums.length == index.length
// 0 <= nums[i] <= 100
// 0 <= index[i] <= i
func createTargetArray(nums []int, indices []int) []int {
	targetArray := make([]int, len(nums))
	for i, index := range indices {
		copy(targetArray[index+1:i+1], targetArray[index:i])
		targetArray[index] = nums[i]
	}
	return targetArray
}
