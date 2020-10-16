// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(numSubmat([][]int{{0, 1, 1, 0}, {0, 1, 1, 1}, {1, 1, 1, 0}}))
	pp.Println(24)
	pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
}

const MaxInt = int(^uint(0) >> 1)

func numSubmat(mat [][]int) int {
	cusum := make([][]int, len(mat))
	for i := range cusum {
		cusum[i] = make([]int, len(mat[0]))
	}

	// for y := range mat {
	// 	for x := range mat[y] {
	// 		if mat[y][x] == 0 {
	// 			continue
	// 		}
	// 		if y == 0 {
	// 			cusum[y][x] = 1
	// 		} else {
	// 			cusum[y][x] = cusum[y-1][x] + 1
	// 		}
	// 	}
	// }

	for y := len(mat) - 1; y >= 0; y-- {
		for x := range mat[y] {
			if mat[y][x] == 0 {
				continue
			}
			if y == len(mat)-1 {
				cusum[y][x] = 1
			} else {
				cusum[y][x] = cusum[y+1][x] + 1
			}
		}
	}

	// pp.Println("=== DEBUG START ======================================")
	// for y := range cusum {
	// 	for x := range cusum[y] {
	// 		fmt.Print(cusum[y][x])
	// 	}
	// 	fmt.Println()
	// }
	// pp.Println("=== DEBUG END ========================================")

	var count int
	for y := range mat {
		for x := range mat[y] {
			height := MaxInt
			for w := 0; w < len(mat[0])-x; w++ {
				if mat[y][x+w] == 0 {
					break
				}
				// pp.Println("=== DEBUG START ======================================")
				// fmt.Printf("y: %d, x: %d\n", y, x)
				// fmt.Printf("width: %d, height: %d, current: %d\n", w, height, cusum[y][x+w])
				// pp.Println("=== DEBUG END ========================================")
				height = min(height, cusum[y][x+w])
				count += height
			}
		}
	}
	return count
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// func numSubmat(mat [][]int) int {
// 	matrix := make([][][2]int, len(mat))
// 	for i := range matrix {
// 		matrix[i] = make([][2]int, len(mat[0]))
// 	}
// 	for y := len(mat) - 1; y >= 0; y-- {
// 		for x := len(mat[0]) - 1; x >= 0; x-- {
// 			if mat[y][x] == 0 {
// 				continue
// 			}

// 			// hor
// 			if x == len(mat[0])-1 {
// 				matrix[y][x][0] = mat[y][x]
// 			} else {
// 				matrix[y][x][0] = matrix[y][x+1][0] + mat[y][x]
// 			}

// 			// ver
// 			if y == len(mat)-1 {
// 				matrix[y][x][1] = mat[y][x]
// 			} else {
// 				matrix[y][x][1] = matrix[y+1][x][1] + mat[y][x]
// 			}
// 		}
// 	}

// 	// pp.Println("=== DEBUG START ======================================")
// 	// for y := range matrix {
// 	// 	for x := range matrix[y] {
// 	// 		fmt.Print(matrix[y][x][0])
// 	// 	}
// 	// 	fmt.Println()
// 	// }
// 	// pp.Println("---")
// 	// for y := range matrix {
// 	// 	for x := range matrix[y] {
// 	// 		fmt.Print(matrix[y][x][1])
// 	// 	}
// 	// 	fmt.Println()
// 	// }
// 	// pp.Println("=== DEBUG END ========================================")

// 	var count int
// 	for y := range matrix {
// 		for x := range matrix[y] {
// 			height, width  := len(mat)-y,len(mat[y])-x
// 			for size := 0; size <= possibility; size++ {
// 			}
// 		}
// 	}
// 	return count
// }

// MaxInt
