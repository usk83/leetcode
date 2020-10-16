// 540. Single Element in a Sorted Array
// https://leetcode.com/problems/single-element-in-a-sorted-array/
package main

import (
	"sort"

	"github.com/k0kubun/pp"
)

// May LeetCoding Challenge Day 12
// https://leetcode.com/explore/featured/card/may-leetcoding-challenge/535/week-2-may-8th-may-14th/3327/
func main() {
	pp.Println("=========================================")
	pp.Println(singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8}))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(singleNonDuplicate([]int{3, 3, 7, 7, 10, 11, 11}))
	pp.Println(10)
	pp.Println("=========================================")
	pp.Println(singleNonDuplicate([]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10}))
	pp.Println(10)
	pp.Println("=========================================")
}

func singleNonDuplicate(nums []int) int {
	return nums[sort.Search(len(nums)>>1, func(i int) bool {
		return nums[i<<1] != nums[i<<1+1]
	})<<1]
}
