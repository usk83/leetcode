// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(minInteger("9438957234785635408", 23))
	pp.Println("0345989723478563548")
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

func minInteger(numString string, k int) string {
	indices := [10][]int{}
	for i := range numString {
		num := int(numString[i] - '0')
		indices[num] = append(indices[num], i)
	}

	// pp.Println("=== DEBUG START ======================================")
	// for _, row := range indices {
	// 	fmt.Println(row)
	// }
	// pp.Println("=== DEBUG END ========================================")

	minIntString := make([]byte, len(numString))
	for i := range minIntString {
		var index int
		for n := range indices {
			if len(indices[n]) != 0 && indices[n][0] <= i+k {
				minIntString[i] = byte(n + '0')
				index = indices[n][0]
				k -= index - i
				indices[n] = indices[n][1:]
				break
			}
		}
		// pp.Println(i, k, index)
		for n := range indices {
			for j := range indices[n] {
				if indices[n][j] > index {
					break
				}
				indices[n][j]++
			}
		}
		// pp.Println("=== DEBUG START ======================================")
		// for _, row := range indices {
		// 	fmt.Println(row)
		// }
		// pp.Println("=== DEBUG END ========================================")
	}
	return string(minIntString)
}
