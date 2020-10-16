// 5424. Maximum Product of Two Elements in an Array
// https://leetcode.com/contest/weekly-contest-191/problems/maximum-product-of-two-elements-in-an-array/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(maxProduct([]int{3, 4, 5, 2}))
	pp.Println(12)
	pp.Println("=========================================")
	pp.Println(maxProduct([]int{1, 5, 4, 5}))
	pp.Println(16)
	pp.Println("=========================================")
	pp.Println(maxProduct([]int{3, 7}))
	pp.Println(12)
	pp.Println("=========================================")

}

/*
 * 2. Filter to keep two biggest numbers
 */
func maxProduct(nums []int) int {
	bigs := [2]int{}
	for _, num := range nums {
		for i := range bigs {
			if bigs[i] < num {
				bigs[i], num = num, bigs[i]
			}
		}
	}
	return (bigs[0] - 1) * (bigs[1] - 1)
}

/*
 * 1. Sort
 */
// func maxProduct(nums []int) int {
// 	sort.Ints(nums)
// 	one, two := nums[len(nums)-1], nums[len(nums)-2]
// 	return (one - 1) * (two - 1)
// }
