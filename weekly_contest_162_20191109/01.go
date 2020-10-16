package main

import "github.com/k0kubun/pp"

func main() {
	pp.Println("=========================================")
	pp.Println(oddCells(2, 3, [][]int{{0, 1}, {1, 1}}))
	pp.Println(6)
	pp.Println("=========================================")
	pp.Println(oddCells(2, 2, [][]int{{1, 1}, {0, 0}}))
	pp.Println(0)
	pp.Println("=========================================")
}

// Given n and m which are the dimensions of a matrix initialized by zeros and given an array indices where indices[i] = [ri, ci]. For each pair of [ri, ci] you have to increment all cells in row ri and column ci by 1.

// Return the number of cells with odd values in the matrix after applying the increment to all indices.

// [0, 1] [1, 1]

// 1, 2, 1
// 1, 2, 1

// [1, 2, 1]
// [0, 1, 0]

// [1, 2, 1]
// [0, 1, 0]

// [1, 3, 1]
// [1, 3, 1]

// 1 <= n <= 50
// 1 <= m <= 50
// 1 <= indices.length <= 100
// 0 <= indices[i][0] < n
// 0 <= indices[i][1] < m
func oddCells(n int, m int, indices [][]int) int {
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, m)
	}

	for _, index := range indices {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if i == index[0] {
					arr[i][j]++
				}
				if j == index[1] {
					arr[i][j]++
				}
			}
		}
	}

	var odds int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			odds += arr[i][j] & 1
		}
	}

	return odds
}
