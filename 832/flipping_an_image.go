// 832. Flipping an Image
// https://leetcode.com/problems/flipping-an-image/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	fmt.Println(flipAndInvertImage([][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}))
	pp.Printf("Expected: ")
	fmt.Println([][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}})

	fmt.Println(flipAndInvertImage([][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}))
	pp.Printf("Expected: ")
	fmt.Println([][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}})
}

func flipAndInvertImage(A [][]int) [][]int {
	yLength, xLength := len(A), len(A[0])
	xMid := xLength / 2
	modFlag := xLength%2 == 1

	for i := 0; i < yLength; i++ {
		for j := 0; j < xMid; j++ {
			A[i][j], A[i][xLength-1-j] = A[i][xLength-1-j]^1, A[i][j]^1
		}
		if modFlag {
			A[i][xMid] = A[i][xMid] ^ 1
		}
	}

	return A
}
