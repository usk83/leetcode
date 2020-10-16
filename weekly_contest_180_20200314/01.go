// 5356. Lucky Numbers in a Matrix
// https://leetcode.com/contest/weekly-contest-180/problems/lucky-numbers-in-a-matrix/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	fmt.Println(luckyNumbers([][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}))
	fmt.Println([]int{15})
	pp.Println("=========================================")
	fmt.Println(luckyNumbers([][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}}))
	fmt.Println([]int{12})
	pp.Println("=========================================")
	fmt.Println(luckyNumbers([][]int{{7, 8}, {1, 2}}))
	fmt.Println([]int{7})
	pp.Println("=========================================")
}

// m == mat.length
// n == mat[i].length
// 1 <= n, m <= 50
// 1 <= matrix[i][j] <= 10^5.
// All elements in the matrix are distinct.
func luckyNumbers(matrix [][]int) []int {
	arr := []int{}
	for y, row := range matrix {
		minX, minNum := 100000000000000, 100000000000000
		for x, num := range row {
			if num < minNum {
				minX = x
				minNum = num
			}
		}
		maxY, maxNum := 0, 0
		for yy, rowrow := range matrix {
			if rowrow[minX] > maxNum {
				maxY = yy
				maxNum = rowrow[minX]
			}
		}

		if y == maxY {
			arr = append(arr, maxNum)
		}
	}
	return arr
}
