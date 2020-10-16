// 5277. Count Square Submatrices with All Ones
// https://leetcode.com/contest/weekly-contest-165/problems/count-square-submatrices-with-all-ones/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(countSquares([][]int{
		{0, 1, 1, 1},
		{1, 1, 1, 1},
		{0, 1, 1, 1},
	}))
	pp.Println(15)
	pp.Println("=========================================")
	pp.Println(countSquares([][]int{
		{1, 0, 1},
		{1, 1, 0},
		{1, 1, 0},
	}))
	pp.Println(7)
	pp.Println("=========================================")
}

// 1 <= arr.length <= 300
// 1 <= arr[0].length <= 300
// 0 <= arr[i][j] <= 1
// 1 <= arr.length <= 300
// 1 <= arr[0].length <= 300
// 0 <= arr[i][j] <= 1
func countSquares(matrix [][]int) int {
	calc := make([][][2]int, len(matrix))
	for i := range matrix {
		calc[i] = make([][2]int, len(matrix[0]))
	}

	for i := len(matrix) - 1; i >= 0; i-- {
		for j := len(matrix[0]) - 1; j >= 0; j-- {
			if matrix[i][j] == 0 {
				continue
			}
			right := 1
			down := 1
			if j != len(matrix[0])-1 {
				right += calc[i][j+1][0]
			}
			if i != len(matrix)-1 {
				down += calc[i+1][j][1]
			}
			calc[i][j] = [2]int{right, down}
		}
	}

	fmt.Println(calc)

	// 今度は左上から見ていく
	// 有効な範囲内で

	return 0
}

// func countSquares(matrix [][]int) int {
// 	// 縦と横の最長を調べる
// 	// 最長以下で調べていく
// 	// 短いほうが優先

// 	// その連続いくつあるか？がその長さあるか？

// 	holi := make([]int, len(matrix[0]))
// 	for i := len(matrix) - 1; i >= 0; i-- {
// 		holiCount := 0
// 		for j := len(matrix[0]) - 1; j >= 0; j-- {
// 			holiCount += matrix[i][j]
// 		}
// 		holi[holiCount-1]++
// 	}

// 	vert := make([]int, len(matrix))
// 	for i := len(matrix[0]) - 1; i >= 0; i-- {
// 		vertCount := 0
// 		for j := len(matrix) - 1; j >= 0; j-- {
// 			vertCount += matrix[j][i]
// 		}
// 		vert[vertCount-1]++
// 	}

// 	fmt.Println(vert)
// 	fmt.Println(holi)

// 	return 0
// }
