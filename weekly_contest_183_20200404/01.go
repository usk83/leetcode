// 5376. Minimum Subsequence in Non-Increasing Order
// https://leetcode.com/problems/xxx/
package main

import (
	"sort"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
}

func minSubsequence(nums []int) []int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	sort.Ints(nums)
	var subSum int
	ans := []int{}
	for i := len(nums) - 1; i >= 0; i-- {
		ans = append(ans, nums[i])
		subSum += nums[i]
		if subSum > sum/2 {
			break
		}
	}
	return ans
}
