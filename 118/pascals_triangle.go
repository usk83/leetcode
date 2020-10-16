// 118. Pascal's Triangle
// https://leetcode.com/problems/pascals-triangle/
package main

import "fmt"

func main() {
	fmt.Println(generate(5))
	fmt.Println([][]int{
		[]int{1},
		[]int{1, 1},
		[]int{1, 2, 1},
		[]int{1, 3, 3, 1},
		[]int{1, 4, 6, 4, 1},
	})
}

func generate(numRows int) [][]int {
	if numRows <= 0 {
		return [][]int{}
	}
	triangle := make([][]int, numRows)
	triangle[0] = []int{1}
	for i := 1; i < numRows; i++ {
		triangle[i] = make([]int, i+1)
		for j := 0; j < i; j++ {
			triangle[i][j] += triangle[i-1][j]
			triangle[i][j+1] += triangle[i-1][j]
		}
	}
	return triangle
}

// func generateV1(numRows int) [][]int {
// 	if numRows <= 0 {
// 		return [][]int{}
// 	}

// 	triangle := make([][]int, numRows)
// 	triangle[0] = []int{1}

// 	for i := 1; i < numRows; i++ {
// 		row := make([]int, i+1)
// 		for j, lim := 0, i+1; j < lim; j++ {
// 			if j == 0 {
// 				row[j] = triangle[i-1][0]
// 				continue
// 			}
// 			if j == lim-1 {
// 				row[j] = triangle[i-1][lim-2]
// 				continue
// 			}
// 			row[j] = triangle[i-1][j-1] + triangle[i-1][j]
// 		}
// 		triangle[i] = row
// 	}
// 	return triangle
// }
