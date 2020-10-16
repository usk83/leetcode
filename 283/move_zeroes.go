// 283. Move Zeroes
// https://leetcode.com/problems/move-zeroes/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	var nums []int
	pp.Println("=========================================")
	nums = []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
	fmt.Println([]int{1, 3, 12, 0, 0})
	pp.Println("=========================================")
	nums = []int{0, 0, 0, 0, 0, 0, 0, 0}
	moveZeroes(nums)
	fmt.Println(nums)
	fmt.Println([]int{0, 0, 0, 0, 0, 0, 0, 0})
	pp.Println("=========================================")
	nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	moveZeroes(nums)
	fmt.Println(nums)
	fmt.Println([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	pp.Println("=========================================")
	nums = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	moveZeroes(nums)
	fmt.Println(nums)
	fmt.Println([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
	pp.Println("=========================================")
	nums = []int{0, 1, 2, 0, 0, 3, 4, 0, 5, 6, 0, 7, 8, 0, 0, 9, 0, 0}
	moveZeroes(nums)
	fmt.Println(nums)
	fmt.Println([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	pp.Println("=========================================")
}

/*
 * the shorter version
 */
func moveZeroes(nums []int) {
	for i, pointer := 0, 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[pointer] = nums[pointer], nums[i]
			pointer++
		}
	}
}

/*
 * the first solution
 */
// func moveZeroes(nums []int) {
// 	var pointer int
// 	for i, num := range nums {
// 		if num == 0 {
// 			continue
// 		}
// 		if i != pointer {
// 			nums[i], nums[pointer] = nums[pointer], nums[i]
// 		}
// 		pointer++
// 	}
// }
