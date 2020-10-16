// 1277. Count Square Submatrices with All Ones
// https://leetcode.com/problems/count-square-submatrices-with-all-ones/
package main

import (
	"github.com/k0kubun/pp"
)

// May LeetCoding Challenge Day 21
// https://leetcode.com/explore/challenge/card/may-leetcoding-challenge/536/week-3-may-15th-may-21st/3336/
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

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 1 <= arr.length <= 300
// 1 <= arr[0].length <= 300
// 0 <= arr[i][j] <= 1

/*
 * 2. Cumulative sum matrix - O(N^3)
 */

const MaxInt = int(^uint(0) >> 1)

func countSquares(matrix [][]int) int {
	height, width := len(matrix), len(matrix[0])
	cusums := make([][][2]int, height+1) // cusums[y][x] -> [horizontal, vertical]
	for i := range cusums {
		cusums[i] = make([][2]int, width+1)
	}

	// 右下から右方向、下方向の継続数の累積和行列を作成
	for y := height - 1; y >= 0; y-- {
		for x := width - 1; x >= 0; x-- {
			if matrix[y][x] == 1 {
				cusums[y][x][0] = cusums[y][x+1][0] + 1
				cusums[y][x][1] = cusums[y+1][x][1] + 1
			}
		}
	}

	// 全ての点ついて検証
	var count int
	for y := range matrix {
		for x := range matrix[0] {
			possibility := MaxInt
			border := min(len(matrix)-y, len(matrix[0])-x)
			for i := 0; i < border; i++ {
				if possibility <= i {
					break
				}
				possibility = min(possibility, i+min(cusums[y+i][x+i][0], cusums[y+i][x+i][1]))
			}
			count += possibility
		}
	}

	return count
}

/*
 * 1. Brute force - O(N^4)
 */
// func countSquares(matrix [][]int) int {
// 	var count int
// 	for y := range matrix {
// 	LOOP:
// 		for x := range matrix[0] {
// 			border := min(len(matrix)-y, len(matrix[0])-x)
// 			for i := 0; i < border; i++ {
// 				for yy := 0; yy < i; yy++ {
// 					if matrix[y+yy][x+i] != 1 {
// 						continue LOOP
// 					}
// 				}
// 				for xx := 0; xx < i; xx++ {
// 					if matrix[y+i][x+xx] != 1 {
// 						continue LOOP
// 					}
// 				}
// 				if matrix[y+i][x+i] != 1 {
// 					continue LOOP
// 				}
// 				count++
// 			}
// 		}
// 	}
// 	return count
// }
