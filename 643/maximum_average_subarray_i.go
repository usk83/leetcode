// 643. Maximum Average Subarray I
// https://leetcode.com/problems/maximum-average-subarray-i/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
	pp.Println(12.75)
	pp.Println("=========================================")
	pp.Println(findMaxAverage([]int{0, 4, 0, 3, 2}, 1))
	pp.Println(4.0)
	pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
}

// 1. 1 <= k <= n <= 30,000.
// 2. Elements of the given array will be in the range [-10,000, 10,000].
func findMaxAverage(nums []int, k int) float64 {
	var cur int
	for _, num := range nums[:k] {
		cur += num
	}
	// pp.Println(cur)
	// prevIndex := 0
	max := cur
	for i, num := range nums[k:] {
		cur = cur - nums[i] + num
		if cur > max {
			max = cur
		}

		// if nums[prevIndex] < num {
		// 	prevIndex++
		// 	// pp.Println("change", cur)
		// 	max = cur
		// }
	}
	return float64(max) / float64(k)
}

// func findMaxAverage(nums []int, k int) float64 {
// 	if len(nums) == 0 {
// 		return 0
// 	}

// 	// [sum, count]
// 	dp := [2][2]int{}
// 	dp[0][0], dp[0][1] = nums[0], 1
// 	dp[1][0], dp[1][1] = 0, 0

// 	// 取った場合と取らない場合の最適解を集める
// 	for _, num := range nums[1:] {
// 		fmt.Println(dp)

// 		next := [2][2]int{}
// 		better := getBetter(dp)
// 		next[0][0], next[0][1] = better[0]+num, better[1]+1
// 		next[1][0], next[1][1] = num, 1
// 		dp = next
// 	}
// 	fmt.Println(dp)

// 	return math.Max(float64(dp[0][0])/float64(dp[0][1]), float64(dp[1][0])/float64(dp[1][1]))
// }

// func getBetter(val [2][2]int) [2]int {
// 	if float64(val[0][0])/float64(val[0][1]) > float64(val[1][0])/float64(val[1][1]) {
// 		return val[0]
// 	}
// 	return val[1]
// }

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }
