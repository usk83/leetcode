// 287. Find the Duplicate Number
// https://leetcode.com/problems/find-the-duplicate-number/
package main

import (
	"github.com/k0kubun/pp"
)

// June LeetCoding Challenge Day 25
// https://leetcode.com/explore/challenge/card/june-leetcoding-challenge/542/week-4-june-22nd-june-28th/3371/
func main() {
	pp.Println("=========================================")
	pp.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(findDuplicate([]int{3, 1, 3, 4, 2}))
	pp.Println(3)
	pp.Println("=========================================")
}

func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		slow, fast = nums[slow], nums[nums[fast]]
	}
	slow = 0
	for slow != fast {
		slow, fast = nums[slow], nums[fast]
	}
	return slow
}
