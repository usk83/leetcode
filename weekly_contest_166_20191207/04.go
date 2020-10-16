// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
}

// m == mat.length
// n == mat[0].length
// 1 <= m <= 3
// 1 <= n <= 3
// mat[i][j] is 0 or 1.
func minFlips(mat [][]int) int {
	m := mat.length
	n := mat[0].length

	if m == 1 && n == 1 {
		if mat[0][0] == 0 {
			return 0
		} else {
			return -1
		}
	}

	if m == 1 && n == 2 ||
		m == 2 && n == 1 {
	}

	return 0
}
