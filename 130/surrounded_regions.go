// 130. Surrounded Regions
// https://leetcode.com/problems/surrounded-regions/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

// June LeetCoding Challenge Day 17
// https://leetcode.com/explore/featured/card/june-leetcoding-challenge/541/week-3-june-15th-june-21st/3363/
func main() {
	var board [][]byte
	pp.Println("=========================================")
	board = [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	solve(board)
	fmt.Println(board)
	fmt.Println([][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'X', 'X'},
	})
	pp.Println("=========================================")
	board = [][]byte{
		{'O', 'O', 'O'},
		{'O', 'O', 'O'},
		{'O', 'O', 'O'},
	}
	solve(board)
	fmt.Println(board)
	fmt.Println([][]byte{
		{'O', 'O', 'O'},
		{'O', 'O', 'O'},
		{'O', 'O', 'O'},
	})
	pp.Println("=========================================")
	board = [][]byte{
		{'X', 'O', 'X'},
		{'X', 'O', 'X'},
		{'X', 'O', 'X'},
	}
	solve(board)
	fmt.Println(board)
	fmt.Println([][]byte{
		{'X', 'O', 'X'},
		{'X', 'O', 'X'},
		{'X', 'O', 'X'},
	})
	pp.Println("=========================================")
	board = [][]byte{
		{'O', 'X', 'O', 'O', 'O', 'O', 'O', 'O', 'O'},
		{'O', 'O', 'O', 'X', 'O', 'O', 'O', 'O', 'X'},
		{'O', 'X', 'O', 'X', 'O', 'O', 'O', 'O', 'X'},
		{'O', 'O', 'O', 'O', 'X', 'O', 'O', 'O', 'O'},
		{'X', 'O', 'O', 'O', 'O', 'O', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'O', 'X', 'O', 'X', 'O', 'X'},
		{'O', 'O', 'O', 'X', 'O', 'O', 'O', 'O', 'O'},
		{'O', 'O', 'O', 'X', 'O', 'O', 'O', 'O', 'O'},
		{'O', 'O', 'O', 'O', 'O', 'X', 'X', 'O', 'O'},
	}
	solve(board)
	fmt.Println(board)
	fmt.Println([][]byte{
		{'O', 'X', 'O', 'O', 'O', 'O', 'O', 'O', 'O'},
		{'O', 'O', 'O', 'X', 'O', 'O', 'O', 'O', 'X'},
		{'O', 'X', 'O', 'X', 'O', 'O', 'O', 'O', 'X'},
		{'O', 'O', 'O', 'O', 'X', 'O', 'O', 'O', 'O'},
		{'X', 'O', 'O', 'O', 'O', 'O', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'O', 'X', 'O', 'X', 'O', 'X'},
		{'O', 'O', 'O', 'X', 'O', 'O', 'O', 'O', 'O'},
		{'O', 'O', 'O', 'X', 'O', 'O', 'O', 'O', 'O'},
		{'O', 'O', 'O', 'O', 'O', 'X', 'X', 'O', 'O'},
	})
	pp.Println("=========================================")
}

// メモリ取って塗る方法

// メモリ取らない方法
func solve(board [][]byte) {
	if len(board) <= 2 || len(board[0]) <= 2 {
		return
	}

	var isSurroundedByX func(x, y int) bool
	isSurroundedByX = func(x, y int) bool {
		if x < 0 || len(board[0]) == x || y < 0 || len(board) == y {
			return false
		}
		if board[y][x] == 'X' || board[y][x] == '!' {
			return true
		}
		board[y][x] = '!'
		yes := isSurroundedByX(x, y-1)
		yes = isSurroundedByX(x+1, y) && yes
		yes = isSurroundedByX(x, y+1) && yes
		yes = isSurroundedByX(x-1, y) && yes
		return yes
	}

	var capture func(x, y int)
	capture = func(x, y int) {
		if x < 0 || len(board[0]) == x || y < 0 || len(board) == y || board[y][x] == 'X' {
			return
		}
		board[y][x] = 'X'
		capture(x, y-1)
		capture(x+1, y)
		capture(x, y+1)
		capture(x-1, y)
	}

	for y := range board[2:] {
		for x := range board[0][2:] {
			if board[y+1][x+1] == 'O' && isSurroundedByX(x+1, y+1) {
				capture(x+1, y+1)
			}
		}
	}

	for y := range board {
		for x := range board[0] {
			if board[y][x] == '!' {
				board[y][x] = 'O'
			}
		}
	}
}
