// 1222. Queens That Can Attack the King
// https://leetcode.com/problems/queens-that-can-attack-the-king/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(queensAttacktheKing([][]int{{0, 1}, {1, 0}, {4, 0}, {0, 4}, {3, 3}, {2, 4}}, []int{0, 0}))
	fmt.Println([][]int{{0, 1}, {1, 0}, {3, 3}})

	fmt.Println(queensAttacktheKing([][]int{{0, 0}, {1, 1}, {2, 2}, {3, 4}, {3, 5}, {4, 4}, {4, 5}}, []int{3, 3}))
	fmt.Println([][]int{{2, 2}, {3, 4}, {4, 4}})

	fmt.Println(queensAttacktheKing([][]int{{5, 6}, {7, 7}, {2, 1}, {0, 7}, {1, 6}, {5, 1}, {3, 7}, {0, 3}, {4, 0}, {1, 2}, {6, 3}, {5, 0}, {0, 4}, {2, 2}, {1, 1}, {6, 4}, {5, 4}, {0, 0}, {2, 6}, {4, 5}, {5, 2}, {1, 4}, {7, 5}, {2, 3}, {0, 5}, {4, 2}, {1, 0}, {2, 7}, {0, 1}, {4, 6}, {6, 1}, {0, 6}, {4, 3}, {1, 7}}, []int{3, 4}))
	fmt.Println([][]int{{2, 3}, {1, 4}, {1, 6}, {3, 7}, {4, 3}, {5, 4}, {4, 5}})
}

// 1 <= queens.length <= 63
// queens[0].length == 2
// 0 <= queens[i][j] < 8
// king.length == 2
// 0 <= king[0], king[1] < 8
// At most one piece is allowed in a cell.
func queensAttacktheKing(queens [][]int, king []int) [][]int {
	queensMap := make(map[[2]int]struct{}, len(queens))
	for _, queen := range queens {
		queensMap[[2]int{queen[0], queen[1]}] = struct{}{}
	}

	attackableQueens := make([][]int, 0, 8)
	for x := -1; x <= 1; x++ {
	SEARCH:
		for y := -1; y <= 1; y++ {
			k := [2]int{king[0], king[1]}
			for k[0] >= 0 && k[0] <= 7 && k[1] >= 0 && k[1] <= 7 {
				if x == 0 && y == 0 {
					continue SEARCH
				}
				k[0] += x
				k[1] += y
				if _, found := queensMap[k]; found {
					attackableQueens = append(attackableQueens, []int{k[0], k[1]})
					continue SEARCH
				}
			}
		}
	}

	return attackableQueens
}
