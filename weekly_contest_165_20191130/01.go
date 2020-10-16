// 5275. Find Winner on a Tic Tac Toe Game
// https://leetcode.com/problems/find-winner-on-a-tic-tac-toe-game/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(tictactoe([][]int{{0, 0}, {2, 0}, {1, 1}, {2, 1}, {2, 2}}))
	pp.Println("A")
	pp.Println("=========================================")
	pp.Println(tictactoe([][]int{{0, 0}, {1, 1}, {0, 1}, {0, 2}, {1, 0}, {2, 0}}))
	pp.Println("B")
	pp.Println("=========================================")
	pp.Println(tictactoe([][]int{{0, 0}, {1, 1}, {2, 0}, {1, 0}, {1, 2}, {2, 1}, {0, 1}, {0, 2}, {2, 2}}))
	pp.Println("Draw")
	pp.Println("=========================================")
	pp.Println(tictactoe([][]int{{0, 0}, {1, 1}}))
	pp.Println("Pending")
	pp.Println("=========================================")
	pp.Println(tictactoe([][]int{{2, 0}, {0, 2}, {0, 0}, {2, 2}, {1, 1}, {1, 0}, {2, 1}, {0, 1}}))
	pp.Println("Pending")
	pp.Println("=========================================")
}

// Constraints:
//     1 <= moves.length <= 9
//     moves[i].length == 2
//     0 <= moves[i][j] <= 2
//     There are no repeated elements on moves.
//     moves follow the rules of tic tac toe.
func tictactoe(moves [][]int) string {
	grid := [3][3]int{}

	before := 1
	for _, move := range moves {
		grid[move[0]][move[1]] = before
		switch before {
		case 1:
			before = 2
		case 2:
			before = 1
		}
	}

	// if (9 - len(moves)) == 1 {
	// 	for i := range grid {
	// 		for j := range grid {
	// 			if grid[i][j] == 0 {
	// 				grid[i][j] = 1
	// 			}
	// 		}
	// 	}
	// }

	fmt.Println(grid)

	end := func(this int) string {
		if this == 1 {
			return "A"
		} else if this == 2 {
			return "B"
		}
		return ""
	}

	half := (9 - len(moves)) / 2
	aReft := half + (9-len(moves))%2
	bReft := half

	aReach := 0
	bReach := 0
	for i := 0; i < 3; i++ {
		if grid[0][i] == grid[1][i] && grid[1][i] == grid[2][i] {
			if winner := end(grid[1][i]); winner != "" {
				return winner
			}
		}
	}
	// if aReach >= aReft && aReft != bReft {

	// }

	aReach = 0
	bReach = 0
	for i := 0; i < 3; i++ {
		if grid[i][0] == grid[i][1] && grid[i][1] == grid[i][2] {
			if winner := end(grid[i][1]); winner != "" {
				return winner
			}
		}
	}

	aReach = 0
	bReach = 0
	if grid[0][0] == grid[1][1] && grid[1][1] == grid[2][2] {
		if winner := end(grid[1][1]); winner != "" {
			return winner
		}
	}

	aReach = 0
	bReach = 0
	if grid[0][2] == grid[1][1] && grid[1][1] == grid[2][0] {
		if winner := end(grid[1][1]); winner != "" {
			return winner
		}
	}

	// if len(moves) >= 8 {
	if len(moves) == 9 {
		return "Draw"
	}

	// あと何回打てるか？
	// 自分の回数より多いとそっちの勝ち

	_ = aReft
	_ = bReft
	_ = aReach
	_ = bReach

	return "Pending"
}
