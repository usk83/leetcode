// 999. Available Captures for Rook
// https://leetcode.com/problems/available-captures-for-rook/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(numRookCaptures([][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', 'R', '.', '.', '.', 'p'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
	}))
	pp.Println(3)
	pp.Println("=========================================")
	pp.Println(numRookCaptures([][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', 'p', 'p', 'p', 'p', 'p', '.', '.'},
		{'.', 'p', 'p', 'B', 'p', 'p', '.', '.'},
		{'.', 'p', 'B', 'R', 'B', 'p', '.', '.'},
		{'.', 'p', 'p', 'B', 'p', 'p', '.', '.'},
		{'.', 'p', 'p', 'p', 'p', 'p', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
	}))
	pp.Println(0)
	pp.Println("=========================================")
	pp.Println(numRookCaptures([][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'p', 'p', '.', 'R', '.', 'p', 'B', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'B', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
	}))
	pp.Println(3)
	pp.Println("=========================================")
}

/*
 * Time: 08:51
 * Runtime: 0 ms, faster than 100.00%
 * Memory Usage: 1.9 MB, less than 100.00%
 */
func numRookCaptures(board [][]byte) int {
	var y, x int
FIND_ROOK:
	for y = range board {
		for x = range board[y] {
			if board[y][x] == 'R' {
				break FIND_ROOK
			}
		}
	}
	var count int
	for _, dir := range [][2]int{[2]int{-1, 0}, [2]int{0, 1}, [2]int{1, 0}, [2]int{0, -1}} {
		for yy, xx := y+dir[0], x+dir[1]; yy >= 0 && yy < 8 && xx >= 0 && xx < 8; yy, xx = yy+dir[0], xx+dir[1] {
			if board[yy][xx] == 'B' {
				break
			}
			if board[yy][xx] == 'p' {
				count++
				break
			}
		}
	}
	return count
}
