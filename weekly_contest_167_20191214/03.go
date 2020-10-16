// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(maxSideLength([][]int{{1, 1, 3, 2, 4, 3, 2}, {1, 1, 3, 2, 4, 3, 2}, {1, 1, 3, 2, 4, 3, 2}}, 4))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(maxSideLength([][]int{{2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}}, 1))
	pp.Println(0)
	pp.Println("=========================================")
	pp.Println(maxSideLength([][]int{{1, 1, 1, 1}, {1, 0, 0, 0}, {1, 0, 0, 0}, {1, 0, 0, 0}}, 6))
	pp.Println(3)
	pp.Println("=========================================")
	pp.Println(maxSideLength([][]int{{18, 70}, {61, 1}, {25, 85}, {14, 40}, {11, 96}, {97, 96}, {63, 45}}, 40184))
	pp.Println(2)
	pp.Println("=========================================")

}

// 1 <= m, n <= 300
// m == mat.length
// n == mat[i].length
// 0 <= mat[i][j] <= 10000
// 0 <= threshold <= 10^5
func maxSideLength(mat [][]int, threshold int) int {
	m := len(mat)
	n := len(mat[0])

	startIndices := make([][2]int, 0, m*n)
	for y := range mat {
		for x := range mat[y] {
			startIndices = append(startIndices, [2]int{x, y})
		}
	}
	max := 1
	for {
		nextIndices := [][2]int{}

	LOOP:
		for _, indices := range startIndices {
			x, y := indices[0], indices[1]

			if n-x < max || m-y < max {
				continue
			}

			sum := 0
			for yy := 0; yy < max; yy++ {
				for xx := 0; xx < max; xx++ {
					sum += mat[y+yy][x+xx]
					if sum > threshold {
						continue LOOP
					}
				}
			}

			nextIndices = append(nextIndices, [2]int{x, y})
		}

		if len(nextIndices) == 0 {
			break
		}
		max++
	}

	return max - 1
}
