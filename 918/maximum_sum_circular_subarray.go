// 918. Maximum Sum Circular Subarray
// https://leetcode.com/problems/maximum-sum-circular-subarray/
package main

import (
	"github.com/k0kubun/pp"
)

// May LeetCoding Challenge Day 15
// https://leetcode.com/explore/featured/card/may-leetcoding-challenge/536/week-3-may-15th-may-21st/3330/
func main() {
	pp.Println("=========================================")
	pp.Println(maxSubarraySumCircular([]int{1, -2, 3, -2}))
	pp.Println(3)
	pp.Println("=========================================")
	pp.Println(maxSubarraySumCircular([]int{5, -3, 5}))
	pp.Println(10)
	pp.Println("=========================================")
	pp.Println(maxSubarraySumCircular([]int{3, -1, 2, -1}))
	pp.Println(4)
	pp.Println("=========================================")
	pp.Println(maxSubarraySumCircular([]int{3, -2, 2, -3}))
	pp.Println(3)
	pp.Println("=========================================")
	pp.Println(maxSubarraySumCircular([]int{-2, -3, -1}))
	pp.Println(-1)
	pp.Println("=========================================")
}

// -30000 <= A[i] <= 30000
// 1 <= A.length <= 30000
func maxSubarraySumCircular(arr []int) int {
	sum, minSum, maxSum := arr[0], arr[0], arr[0]
	var curMinSum, curMaxSum int
	if arr[0] < 0 {
		curMinSum = arr[0]
	} else {
		curMaxSum = arr[0]
	}

	for _, num := range arr[1:] {
		sum, curMinSum, curMaxSum = sum+num, curMinSum+num, curMaxSum+num
		minSum, maxSum = min(minSum, curMinSum), max(maxSum, curMaxSum)
		if curMinSum > 0 {
			curMinSum = 0
		}
		if curMaxSum < 0 {
			curMaxSum = 0
		}
	}

	if sum == minSum {
		return maxSum
	}
	return max(maxSum, sum-minSum)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
