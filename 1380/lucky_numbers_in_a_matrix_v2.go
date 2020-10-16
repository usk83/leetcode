// 1380. Lucky Numbers in a Matrix
// https://leetcode.com/problems/lucky-numbers-in-a-matrix/
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

/*
 * straightforward solution with memorization
 * Time Complexity: O(M*N + M*M)
 * Space Complexity: O(N)
 */
func luckyNumbers(matrix [][]int) []int {
	colMaxes := make([]*int, len(matrix[0]))
	for _, row := range matrix {
		xIndex := len(row) - 1
		for x := range row[:xIndex] {
			if row[x] < row[xIndex] {
				xIndex = x
			}
		}
		if colMaxes[xIndex] == nil {
			yIndex := len(matrix) - 1
			for y := range matrix[:yIndex] {
				if matrix[y][xIndex] > matrix[yIndex][xIndex] {
					yIndex = y
				}
			}
			colMaxes[xIndex] = &matrix[yIndex][xIndex]
		}
		if row[xIndex] == *colMaxes[xIndex] {
			return []int{row[xIndex]}
		}
	}
	return nil
}

/*
 * straightforward solution
 * Time Complexity: O(M*N + M*M)
 * Space Complexity: O(1)
 */
func luckyNumbers(matrix [][]int) []int {
	for _, row := range matrix {
		xIndex := len(row) - 1
		for x := range row[:xIndex] {
			if row[x] < row[xIndex] {
				xIndex = x
			}
		}
		yIndex := len(matrix) - 1
		for y := range matrix[:yIndex] {
			if matrix[y][xIndex] > matrix[yIndex][xIndex] {
				yIndex = y
			}
		}
		if row[xIndex] == matrix[yIndex][xIndex] {
			return []int{row[xIndex]}
		}
	}
	return nil
}
