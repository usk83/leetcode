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

func luckyNumbers(matrix [][]int) []int {
	rowMins := make([]*int, len(matrix))
	colMaxes := make([]*int, len(matrix[0]))
	var rowMin func(row int, val *int) (*int, bool)
	var colMax func(col int, val *int) (*int, bool)
	rowMin = func(row int, val *int) (*int, bool) {
		if rowMins[row] != nil {
			return val, val == rowMins[row]
		}
		minIndex := 0
		for index := range matrix[row][1:] {
			if matrix[row][index+1] < matrix[row][minIndex] {
				minIndex = index + 1
			}
		}
		if val == &matrix[row][minIndex] {
			return val, true
		}
		rowMins[row] = &matrix[row][minIndex]
		return colMax(minIndex, &matrix[row][minIndex])
	}
	colMax = func(col int, val *int) (*int, bool) {
		if colMaxes[col] != nil {
			return val, val == colMaxes[col]
		}
		maxIndex := 0
		for index := range matrix[1:] {
			if matrix[index+1][col] > matrix[maxIndex][col] {
				maxIndex = index + 1
			}
		}
		if val == &matrix[maxIndex][col] {
			return val, true
		}
		colMaxes[col] = &matrix[maxIndex][col]
		return rowMin(maxIndex, &matrix[maxIndex][col])
	}

	for row := range matrix {
		if num, luck := rowMin(row, nil); luck {
			return []int{*num}
		}
	}
	return nil
}
