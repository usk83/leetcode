package main

import "fmt"

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
	kingX, kingY := king[1], king[0]
	result := [][]int{}

	// kingから数えていけばいい

	// leftUpPosition
	tmpKing := []int{kingY, kingX}
	loop := true
	for loop {
		tmpKing[0]--
		tmpKing[1]--
		if tmpKing[0] >= 0 && tmpKing[0] <= 7 && tmpKing[1] >= 0 && tmpKing[1] <= 7 {
			for _, queen := range queens {
				if queen[0] == tmpKing[0] && queen[1] == tmpKing[1] {
					result = append(result, queen)
					loop = false
				}
			}
		} else {
			loop = false
		}
	}
	// upPosition
	tmpKing = []int{kingY, kingX}
	loop = true
	for loop {
		tmpKing[0]--
		if tmpKing[0] >= 0 && tmpKing[0] <= 7 && tmpKing[1] >= 0 && tmpKing[1] <= 7 {
			for _, queen := range queens {
				if queen[0] == tmpKing[0] && queen[1] == tmpKing[1] {
					result = append(result, queen)
					loop = false
				}
			}
		} else {
			loop = false
		}
	}
	// rightUpPosition
	tmpKing = []int{kingY, kingX}
	loop = true
	for loop {
		tmpKing[0]--
		tmpKing[1]++
		if tmpKing[0] >= 0 && tmpKing[0] <= 7 && tmpKing[1] >= 0 && tmpKing[1] <= 7 {
			for _, queen := range queens {
				if queen[0] == tmpKing[0] && queen[1] == tmpKing[1] {
					result = append(result, queen)
					loop = false
				}
			}
		} else {
			loop = false
		}
	}
	// rightPosition
	tmpKing = []int{kingY, kingX}
	loop = true
	for loop {
		tmpKing[1]++
		if tmpKing[0] >= 0 && tmpKing[0] <= 7 && tmpKing[1] >= 0 && tmpKing[1] <= 7 {
			for _, queen := range queens {
				if queen[0] == tmpKing[0] && queen[1] == tmpKing[1] {
					result = append(result, queen)
					loop = false
				}
			}
		} else {
			loop = false
		}
	}
	// rightDownPosition
	tmpKing = []int{kingY, kingX}
	loop = true
	for loop {
		tmpKing[0]++
		tmpKing[1]++
		if tmpKing[0] >= 0 && tmpKing[0] <= 7 && tmpKing[1] >= 0 && tmpKing[1] <= 7 {
			for _, queen := range queens {
				if queen[0] == tmpKing[0] && queen[1] == tmpKing[1] {
					result = append(result, queen)
					loop = false
				}
			}
		} else {
			loop = false
		}
	}
	// downPosition
	tmpKing = []int{kingY, kingX}
	loop = true
	for loop {
		tmpKing[0]++
		if tmpKing[0] >= 0 && tmpKing[0] <= 7 && tmpKing[1] >= 0 && tmpKing[1] <= 7 {
			for _, queen := range queens {
				if queen[0] == tmpKing[0] && queen[1] == tmpKing[1] {
					result = append(result, queen)
					loop = false
				}
			}
		} else {
			loop = false
		}
	}
	// leftDownPosition
	tmpKing = []int{kingY, kingX}
	loop = true
	for loop {
		tmpKing[0]++
		tmpKing[1]--
		if tmpKing[0] >= 0 && tmpKing[0] <= 7 && tmpKing[1] >= 0 && tmpKing[1] <= 7 {
			for _, queen := range queens {
				if queen[0] == tmpKing[0] && queen[1] == tmpKing[1] {
					result = append(result, queen)
					loop = false
				}
			}
		} else {
			loop = false
		}
	}
	// leftPosition
	tmpKing = []int{kingY, kingX}
	loop = true
	for loop {
		tmpKing[1]--
		if tmpKing[0] >= 0 && tmpKing[0] <= 7 && tmpKing[1] >= 0 && tmpKing[1] <= 7 {
			for _, queen := range queens {
				if queen[0] == tmpKing[0] && queen[1] == tmpKing[1] {
					result = append(result, queen)
					loop = false
				}
			}
		} else {
			loop = false
		}
	}

	return result
}
