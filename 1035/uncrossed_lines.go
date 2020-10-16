// 1035. Uncrossed Lines
// https://leetcode.com/problems/uncrossed-lines/
package main

import (
	"github.com/k0kubun/pp"
)

// May LeetCoding Challenge Day 25
// https://leetcode.com/explore/challenge/card/may-leetcoding-challenge/537/week-4-may-22nd-may-28th/3340/
func main() {
	pp.Println("=========================================")
	pp.Println(maxUncrossedLines([]int{1, 4, 2}, []int{1, 2, 4}))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(maxUncrossedLines([]int{2, 5, 1, 2, 5}, []int{10, 5, 2, 1, 5, 2}))
	pp.Println(3)
	pp.Println("=========================================")
	pp.Println(maxUncrossedLines([]int{1, 3, 7, 1, 7, 5}, []int{1, 9, 2, 5, 1}))
	pp.Println(2)
	pp.Println("=========================================")
}

// 1 <= A.length <= 500
// 1 <= B.length <= 500
// 1 <= A[i], B[i] <= 2000

/*
 * 2. Longest Common Subsequence
 *   - Time Complexity: O(N^2)
 *   - Space Complexity: O(N)
 */
func maxUncrossedLines(top []int, bottom []int) int {
	dp := [2][]int{}
	for i := range dp {
		dp[i] = make([]int, len(bottom)+1)
	}
	for i, n1 := range top {
		for j, n2 := range bottom {
			if n1 == n2 {
				dp[i&1^1][j+1] = dp[i&1][j] + 1
			} else {
				dp[i&1^1][j+1] = max(dp[i&1][j+1], dp[i&1^1][j])
			}
		}
	}
	return dp[len(top)&1][len(bottom)]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
 * 1. DP
 * 上の数字を左からみていく
 * 現時点で同じ本数の線を引けている場合はより左側の点を使っているものを選択する
 * 数字の検索をするために数字ごとのindicesのスライスを作成する
 *   Binary Search で logN で引ける
 *   - Time Complexity: O(N^2 logN)
 *   - Space Complexity: O(N)
 */
// func maxUncrossedLines(top []int, bottom []int) int {
// 	indices := map[int][]int{}
// 	for i, num := range bottom {
// 		indices[num] = append(indices[num], i)
// 	}
//
// 	dp := []int{-1}
// 	for _, num := range top {
// 		if len(indices[num]) == 0 {
// 			continue
// 		}
//
// 		next := make([]int, len(dp), len(dp)+1)
// 		copy(next, dp)
//
// 		for count, usedIndex := range dp {
// 			target := sort.Search(len(indices[num]), func(i int) bool {
// 				return indices[num][i] > usedIndex
// 			})
// 			if target == len(indices[num]) {
// 				continue
// 			}
// 			if count+1 < len(next) {
// 				next[count+1] = min(next[count+1], indices[num][target])
// 			} else {
// 				next = append(next, indices[num][target])
// 			}
// 		}
//
// 		dp = next
// 	}
// 	return len(dp) - 1
// }
//
// func min(x, y int) int {
// 	if x < y {
// 		return x
// 	}
// 	return y
// }
