// 1275. Find Winner on a Tic Tac Toe Game
// https://leetcode.com/problems/find-winner-on-a-tic-tac-toe-game/
package main

import (
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
	for i, move := range moves {
		grid[move[0]][move[1]] = i&1 + 1
	}

	for i := 0; i < 3; i++ {
		if player := grid[i][i]; player != 0 {
			if (grid[i][0] == grid[i][1] && grid[i][0] == grid[i][2]) ||
				(grid[0][i] == grid[1][i] && grid[0][i] == grid[2][i]) {
				return string(byte('@' + player))
			}
		}
	}

	if player := grid[1][1]; player != 0 {
		if (grid[0][0] == grid[2][2] && grid[0][0] == player) ||
			(grid[2][0] == grid[0][2] && grid[2][0] == player) {
			return string(byte('@' + player))
		}
	}

	if len(moves) != 9 {
		return "Pending"
	}

	return "Draw"
}
