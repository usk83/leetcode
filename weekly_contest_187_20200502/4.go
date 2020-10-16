// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
	"sort"
	"strconv"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(kthSmallest([][]int{{1, 3, 11}, {2, 4, 6}}, 5))
	pp.Println(7)
	pp.Println("=========================================")
	pp.Println(kthSmallest([][]int{{1, 3, 11}, {2, 4, 6}}, 9))
	pp.Println(17)
	pp.Println("=========================================")
	pp.Println(kthSmallest([][]int{{1, 10, 10}, {1, 4, 5}, {2, 3, 6}}, 7))
	pp.Println(9)
	pp.Println("=========================================")
	pp.Println(kthSmallest([][]int{{1, 1, 10}, {2, 2, 9}}, 7))
	pp.Println(12)
	pp.Println("=========================================")
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

type state struct {
	sum       int
	selection []int
	used      bool
}

func kthSmallest(matrix [][]int, k int) int {
	n, m := len(matrix[0]), len(matrix)

	diffMatrix := make([][]int, 0, m)
	for _, row := range matrix {
		diffRow := make([]int, n)
		for i := range row {
			if i != len(row)-1 {
				diffRow[i] = row[i+1] - row[i]
			} else {
				diffRow[i] = -1
			}
		}
		diffMatrix = append(diffMatrix, diffRow)
	}

	firstState := state{}
	firstState.selection = make([]int, m)
	for _, row := range matrix {
		firstState.sum += row[0]
	}
	sums := []state{firstState}

	for count := 0; count < k; count++ {
		uniq := map[string]bool{}
		for _, sum := range sums {
			key := ""
			for _, item := range sum.selection {
				key += strconv.Itoa(item) + "-"
			}
			uniq[key] = true
		}

		// pp.Println("=== DEBUG START ======================================")
		// fmt.Println(sums)
		// pp.Println("=== DEBUG END ========================================")
		sums[count].used = true

		nextSums := make([]state, len(sums), len(sums)+m)
		copy(nextSums, sums)

		// 次を作成する
		for rowIndex := 0; rowIndex < m; rowIndex++ {
			nextState := state{}
			nextState.sum = sums[count].sum
			nextState.selection = make([]int, m)
			copy(nextState.selection, sums[count].selection)

			if diffMatrix[rowIndex][nextState.selection[rowIndex]] < 0 {
				continue
			}
			nextState.sum += diffMatrix[rowIndex][nextState.selection[rowIndex]]
			nextState.selection[rowIndex]++

			key := ""
			for _, item := range nextState.selection {
				key += strconv.Itoa(item) + "-"
			}
			// pp.Println("=== DEBUG START ======================================")
			// pp.Println(key)
			// pp.Println("=== DEBUG END ========================================")
			if !uniq[key] {
				nextSums = append(nextSums, nextState)
				uniq[key] = true
			}
		}

		sort.Slice(nextSums, func(i, j int) bool {
			if nextSums[i].used == nextSums[j].used {
				return nextSums[i].sum < nextSums[j].sum
			}
			if nextSums[i].used {
				return true
			}
			return false
		})
		if len(nextSums) <= k {
			sums = nextSums
		} else {
			sums = nextSums[:k]
		}
	}
	return sums[k-1].sum
}
