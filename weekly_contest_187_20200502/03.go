// 5402. Longest Continuous Subarray With Absolute Diff Less Than or Equal to Limit
// https://leetcode.com/contest/weekly-contest-187/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(longestSubarray([]int{8, 2, 4, 7}, 4))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(longestSubarray([]int{10, 1, 2, 4, 7, 2}, 5))
	pp.Println(4)
	pp.Println("=========================================")
	pp.Println(longestSubarray([]int{4, 2, 2, 2, 4, 4, 2, 2}, 0))
	pp.Println(3)
	pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
}

func longestSubarray(nums []int, limit int) int {
	start := 0
	longest := 0
	index := 0
	mn, mx := nums[0], nums[0]
	for i, num := range nums[1:] {
		index = i + 1
		mn = min(mn, num)
		mx = max(mx, num)
		if mx-mn > limit {
			longest = max(longest, index-start)
			// 超えたときの処理
			// 追いかける
			// 逆方向に伸ばせるだけ伸ばす
			mx, mn = num, num
			var newStart int
			for newStart = index - 1; mx-mn <= limit && newStart != start; newStart-- {
				mn = min(mn, nums[newStart])
				mx = max(mx, nums[newStart])
			}
			start = newStart + 1
		}
		// pp.Println("start: %s, index: %s\n", start, index)
	}
	longest = max(longest, index-start+1)
	// if longest == 1 {
	// 	return 0
	// }
	return longest
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// func abs(x, y int) int {
// 	return int(math.Abs(float64(x), float64(y)))
// }
