// 5394. Diagonal Traverse II
// https://leetcode.com/contest/weekly-contest-186/problems/diagonal-traverse-ii/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	fmt.Println(findDiagonalOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println([]int{1, 4, 2, 7, 5, 3, 8, 6, 9})
	pp.Println("=========================================")
	fmt.Println(findDiagonalOrder([][]int{{1, 2, 3, 4, 5}, {6, 7}, {8}, {9, 10, 11}, {12, 13, 14, 15, 16}}))
	fmt.Println([]int{1, 6, 2, 8, 7, 3, 9, 4, 12, 10, 5, 13, 11, 14, 15, 16})
	pp.Println("=========================================")
	fmt.Println(findDiagonalOrder([][]int{{1, 2, 3}, {4}, {5, 6, 7}, {8}, {9, 10, 11}}))
	fmt.Println([]int{1, 4, 2, 5, 3, 8, 6, 9, 7, 10, 11})
	pp.Println("=========================================")
	fmt.Println(findDiagonalOrder([][]int{{1, 2, 3, 4, 5, 6}}))
	fmt.Println([]int{1, 2, 3, 4, 5, 6})
	pp.Println("=========================================")
	fmt.Println(findDiagonalOrder([][]int{{1, 2, 3, 4, 5, 6}, {1}}))
	fmt.Println([]int{1, 1, 2, 3, 4, 5, 6})
	pp.Println("=========================================")
}

type LineInfo struct {
	no   int
	data []int
}

func findDiagonalOrder(nums [][]int) []int {
	length, maxIndex := 0, 0
	matrix := make([]LineInfo, len(nums))
	for no, list := range nums {
		maxIndex = max(maxIndex, len(list)-(len(nums)-1-no))
		length += len(list)
		matrix[no] = LineInfo{
			no:   no,
			data: nums[no],
		}
	}

	reorder := make([]int, 0, length)
	startIndex := 0 - len(nums) + 1
	for i := startIndex; i <= maxIndex; i++ {
		next := make([]LineInfo, len(matrix))
		nextLength := 0

		if len(matrix) == 1 {
			line := len(matrix) - 1 - 0
			index := i + len(nums) - 1 - matrix[line].no
			if index >= 0 && index < len(matrix[line].data) {
				reorder = append(reorder, matrix[line].data[index:]...)
			}
			break
		} else {
			for j := 0; j < len(matrix); j++ {
				line := len(matrix) - 1 - j
				index := i + len(nums) - 1 - matrix[line].no
				if index >= 0 && index < len(matrix[line].data) {
					reorder = append(reorder, matrix[line].data[index])
				}
				if index+1 == len(matrix[line].data) {
					continue
				}
				nextLength++
				next[len(next)-nextLength] = matrix[line]
			}
			matrix = next[len(next)-nextLength:]
		}
	}
	return reorder
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// func findDiagonalOrder(nums [][]int) []int {
// 	length := 0
// 	maxIndex := 0
// 	for no, list := range nums {
// 		maxIndex = max(maxIndex, len(list)-(len(nums)-1-no))
// 		length += len(list)
// 	}

// 	finished := 0
// 	reorder := make([]int, 0, length)
// 	startIndex := 0 - len(nums) + 1
// LOOP:
// 	for i := startIndex; i <= maxIndex; i++ {
// 		for j := 0; j < len(nums); j++ {
// 			lineNo := len(nums) - 1 - j
// 			index := i + j
// 			if index >= 0 && index < len(nums[lineNo]) {
// 				reorder = append(reorder, nums[lineNo][index])
// 				if index+1 == len(nums[lineNo]) {
// 					finished++
// 				}
// 			}
// 		}
// 	}

// 	return reorder
// }

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }
