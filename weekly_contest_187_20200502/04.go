// 5403. Find the Kth Smallest Sum of a Matrix With Sorted Rows
// https://leetcode.com/contest/weekly-contest-187/problems/find-the-kth-smallest-sum-of-a-matrix-with-sorted-rows/
package main

import (
	"fmt"
	"sort"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(kthSmallest([][]int{{1, 3, 11}, {2, 4, 6}}, 5))
	pp.Println(7)
	pp.Println("=========================================")
	// pp.Println(kthSmallest([][]int{{1, 3, 11}, {2, 4, 6}}, 9))
	// pp.Println(17)
	// pp.Println("=========================================")
	// pp.Println(kthSmallest([][]int{{1, 10, 10}, {1, 4, 5}, {2, 3, 6}}, 7))
	// pp.Println(9)
	// pp.Println("=========================================")
	// pp.Println(kthSmallest([][]int{{1, 1, 10}, {2, 2, 9}}, 7))
	// pp.Println(12)
	// pp.Println("=========================================")
}

// m == mat.length
// n == mat.length[i]
// 1 <= m, n <= 40
// 1 <= k <= min(200, n ^ m)
// 1 <= mat[i][j] <= 5000
// mat[i] is a non decreasing array.

const (
	MaxUint = ^uint(0)
	MinUint = 0
	MaxInt  = int(^uint(0) >> 1)
	MinInt  = -MaxInt - 1
)

func kthSmallest(matrix [][]int, k int) int {
	// 行は昇順にソートされている
	// 各行ごとに1つ数字を取る
	// k番目に小さい合計はなにか？

	// 各行ごとdiffの配列に交換する
	// 	maxからの

	// 最大値を求める
	// 各行ごとに次のやつからの差分の配列に変換する
	// 次の数字への差分と行とindexを保持して、差分でソートし続ける
	// 小さい差分をとってk回小さくする

	var sum int
	for _, row := range matrix {
		sum += row[len(row)-1]
	}

	m := len(matrix)
	n := len(matrix[0])
	diffMatrix := make([][]int, 0, m)
	// for i := range diffMatrix {
	// 	diffMatrix[i] = make([]int, n)
	// }
	for _, row := range matrix {
		diffRow := make([]int, n)
		for i := range row {
			if i == 0 {
				diffRow[i] = MaxInt
			} else {
				diffRow[i] = row[i] - row[i-1]
			}

		}
		diffMatrix = append(diffMatrix, diffRow)
	}

	diffs := make([][3]int, 0, m)
	// diff, row, index
	for i, row := range matrix {
		diffs = append(diffs, [3]int{diffMatrix[i][len(row)-1], i, len(row) - 1})
	}

	sort.Slice(diffs, func(i, j int) bool {
		return diffs[i][0] < diffs[j][0]
	})

	for count := 0; count < k; count++ {
		sort.Slice(diffs, func(i, j int) bool {
			return diffs[i][0] < diffs[j][0]
		})

		pp.Println("=== DEBUG START ======================================")
		fmt.Println(diffMatrix)
		fmt.Println(diffs)
		// pp.Println(diffs[0][1])
		// pp.Println(diffs[0][2])
		pp.Println("=== DEBUG END ========================================")
		sum -= diffs[0][0]
		// diffs[0][1]
		diffs[0][2]--

		diffs[0][0] = diffMatrix[diffs[0][1]][diffs[0][2]]
	}

	// pp.Println("=== DEBUG START ======================================")
	// pp.Println(sum)
	// fmt.Println(diffMatrix)
	// fmt.Println(diffs)
	// pp.Println("=== DEBUG END ========================================")

	return sum
}
