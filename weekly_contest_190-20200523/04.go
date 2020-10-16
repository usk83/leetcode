// 5419. Max Dot Product of Two Subsequences
// https://leetcode.com/contest/weekly-contest-190/problems/max-dot-product-of-two-subsequences/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(maxDotProduct([]int{2, 1, -2, 5}, []int{3, 0, -6}))
	pp.Println(18)
	pp.Println("=========================================")
	pp.Println(maxDotProduct([]int{3, -2}, []int{2, -6, 7}))
	pp.Println(21)
	pp.Println("=========================================")
	pp.Println(maxDotProduct([]int{-1, -1}, []int{1, 1}))
	pp.Println(-1)
	pp.Println("=========================================")
	pp.Println(maxDotProduct([]int{5, -4, -3}, []int{-4, -3, 0, -4, 2}))
	pp.Println(28)
	pp.Println("=========================================")
}

const (
	MaxUint = ^uint(0)
	MinUint = 0
	MaxInt  = int(^uint(0) >> 1)
	MinInt  = -MaxInt - 1
)

func maxDotProduct(nums1 []int, nums2 []int) int {
	// index1 と index2 の組み合わせごとに最大ケースのみ保持
	// 	2次元配列
	// index1側でループ

	dp := make([][]int, len(nums1)+1)
	for i := range dp {
		dp[i] = make([]int, len(nums2)+1)
		for j := range dp[i] {
			dp[i][j] = MinInt
		}
	}

	// pp.Println("=== DEBUG START ======================================")
	// for _, row := range dp {
	// 	fmt.Println(row)
	// }
	// pp.Println("=== DEBUG END ========================================")

	maxDotProduct := MinInt
	dp[1][0] = 0
	for index2 := range dp[0][1:] {
		current := nums1[0] * nums2[index2]
		dp[1][index2+1] = current
		maxDotProduct = max(maxDotProduct, current)
	}

	// pp.Println("=== DEBUG START ======================================")
	// for _, row := range dp {
	// 	fmt.Println(row)
	// }
	// pp.Println("=== DEBUG END ========================================")

	for i := range dp[1:] {
		// index1 := i + 1
		index1 := i

		for index2 := range dp[index1] {
			current := dp[index1][index2]

			// 選ばない
			// dp[index1+1][index2] = max(dp[index1+1][index2], current)
			if dp[index1+1][index2] == MinInt {
				dp[index1+1][index2] = current
			} else {
				dp[index1+1][index2] = max(dp[index1+1][index2], current)
			}

			if index2 == len(nums2) {
				continue
			}

			if current < 0 {
				current = 0
			}

			// 選ぶ
			// for i2 := index2; i2 < len(nums2); i2++ {
			// current = dp[index1][i2]
			i2 := index2
			next := current + nums1[index1]*nums2[i2]

			// pp.Println("index1:", index1, " index2:", i2, " before:", current, " after:", next)
			// if index1 == 1 && i2 == 1 {
			// 	pp.Println("debugging...")
			// 	pp.Println(dp[index1+1][i2+1])
			// }

			if dp[index1+1][i2+1] == MinInt {
				if index1 == 1 && i2 == 1 {
					dp[index1+1][i2+1] = next
					// pp.Println("yes")
					// pp.Println("=== DEBUG START ======================================")
					// for _, row := range dp {
					// 	fmt.Println(row)
					// }
					// pp.Println("=== DEBUG END ========================================")
				}
			} else {
				dp[index1+1][i2+1] = max(dp[index1+1][i2+1], next)
			}

			// dp[index1+1][i2+1] = next
			maxDotProduct = max(maxDotProduct, next)

			// pp.Println("=== DEBUG START ======================================")
			// pp.Println(maxDotProduct)
			// pp.Println("=== DEBUG END ========================================")
			// }
		}
		// pp.Println("=== DEBUG START ======================================")
		// for _, row := range dp {
		// 	fmt.Println(row)
		// }
		// pp.Println("=== DEBUG END ========================================")
	}

	pp.Println("=== DEBUG START ======================================")
	for _, row := range dp {
		fmt.Println(row)
	}
	pp.Println("=== DEBUG END ========================================")

	return maxDotProduct
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// for index1 := range dp {
// 	if index1+1 == len(nums1) {
// 		break
// 	}
// 	for index2 := range dp[index1] {
// 		current := dp[index1][index2]
// 		maxDotProduct = max(maxDotProduct, current)

// 		// 選ばない
// 		dp[index1+1][index2] = max(dp[index1+1][index2], current)

// 		// 選ぶ
// 		for i := index2; i < len(nums2); i++ {
// 			// bk(index1+1, i+1, current+nums1[index1]*nums2[i])

// 			dp[index1+1][i] = max(dp[index1+1][i], current+nums1[index1]*nums2[i])
// 		}
// 	}
// }

// return maxDotProduct
// }

// 1 <= nums1.length, nums2.length <= 500
// -1000 <= nums1[i], nums2[i] <= 1000
// func _maxDotProduct(nums1 []int, nums2 []int) int {
// 	// maxDotProduct, current := MinInt, 0
// 	// for _, n1 := range nums1 {
// 	// 	for _, n2 := range nums2 {
// 	// 	}
// 	// }
// 	// return maxDotProduct

// 	maxDotProduct := MinInt
// 	var bk func(index1, index2, current int)
// 	bk = func(index1, index2, current int) {
// 		if index1 >= len(nums1) || index2 >= len(nums2) {
// 			return
// 		}

// 		// 選ぶ
// 		for i := index2; i < len(nums2); i++ {
// 			bk(index1+1, i+1, current+nums1[index1]*nums2[i])
// 		}

// 		// 選ばない
// 		bk(index1 + 1)
// 	}
// 	bk(0, 0, 0)
// 	return maxDotProduct
// }

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }
