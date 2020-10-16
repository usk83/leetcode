// 1252. Cells with Odd Values in a Matrix
// https://leetcode.com/problems/cells-with-odd-values-in-a-matrix/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
}

/*
 * v2. tracking a state of each rows and cols and calculating a result
 */
func oddCells(n int, m int, indices [][]int) int {
	row, col := make([]int, n), make([]int, m)
	for _, index := range indices {
		row[index[0]] ^= 1
		col[index[1]] ^= 1
	}

	var rowOdd, colOdd int
	for _, r := range row {
		rowOdd += r
	}
	for _, c := range col {
		colOdd += c
	}

	return rowOdd*(m-colOdd) + (n-rowOdd)*colOdd
}

/*
 * v1. calculating all values in matrix and counting
 */
// func oddCells(n int, m int, indices [][]int) int {
// 	matrix := make([][]int, n)
// 	for i := range matrix {
// 		matrix[i] = make([]int, m)
// 	}
//
// 	for _, pair := range indices {
// 		for i := range matrix[0] {
// 			matrix[pair[0]][i]++
// 		}
// 		for i := range matrix {
// 			matrix[i][pair[1]]++
// 		}
// 	}
//
// 	var count int
// 	for row := range matrix {
// 		for num := range row {
// 			count += num & 1
// 		}
// 	}
// 	return count
// }
