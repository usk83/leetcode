// 5500. Maximum Length of Subarray With Positive Product
// https://leetcode.com/contest/weekly-contest-204/problems/maximum-length-of-subarray-with-positive-product/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(getMaxLen([]int{1, -2, -3, 4}))
	pp.Println(4)
	pp.Println("=========================================")
	pp.Println(getMaxLen([]int{0, 1, -2, -3, -4}))
	pp.Println(3)
	pp.Println("=========================================")
	pp.Println(getMaxLen([]int{-1, -2, -3, 0, 1}))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(getMaxLen([]int{-1, 2}))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(getMaxLen([]int{1, 2, 3, 5, -6, 4, 0, 10}))
	pp.Println(4)
	pp.Println("=========================================")
}

// 0を含んではいけない
// マイナスなら整数個

// ゼロの位置を探して
// その間のマイナスの値のindexと数を確保しておく
// 整数個なら問題なし
// 奇数個なら長い方を考える

func getMaxLen(nums []int) int {
	maxLen := 0
	start := 0
	left, right := -1, -1
	count := 0
	for i := start; i <= len(nums); i++ {
		if i == len(nums) || nums[i] == 0 {
			// pp.Println("確定:", i)
			if count%2 == 0 {
				// pp.Println("偶数")
				maxLen = max(maxLen, i-start)
			} else {
				// pp.Println("奇数")
				rightLen := i - left - 1
				leftLen := right - start
				// pp.Println("start:", start)
				// pp.Println("right:", right)
				// pp.Println("left:", left)
				// pp.Println("right:", rightLen)
				// pp.Println("left:", leftLen)
				maxLen = max(maxLen, max(rightLen, leftLen))
			}
			// pp.Println("暫定:", maxLen)
			start = i + 1
			left, right = -1, -1
			count = 0
		} else {
			if nums[i] < 0 {
				count++
				if left < 0 {
					left, right = i, i
				} else {
					right = i
				}
			}
		}
	}
	return maxLen
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
