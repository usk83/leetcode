// 53. Maximum Subarray
// https://leetcode.com/problems/maximum-subarray/
package main

import "github.com/k0kubun/pp"

func main() {
	pp.Println("=========================================")
	pp.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	pp.Println(6)
	pp.Println("=========================================")
	pp.Println(maxSubArray([]int{-1}))
	pp.Println(-1)
	pp.Println("=========================================")
	pp.Println(maxSubArray([]int{1, 2}))
	pp.Println(3)
	pp.Println("=========================================")
}

func maxSubArray(nums []int) int {
	maxSum, currentSum := nums[0], nums[0]
	for _, num := range nums[1:] {
		if currentSum += num; num > currentSum {
			currentSum = num
		}
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum
}
