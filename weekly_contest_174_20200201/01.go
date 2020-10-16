// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
	"sort"

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

func kWeakestRows(mat [][]int, k int) []int {
	counts := [][2]int{}
	for i, m := range mat {
		count := 0
		for _, s := range m {
			if s == 1 {
				count++
			} else {
				break
			}
		}
		counts = append(counts, [2]int{i, count})
	}

	sort.Slice(counts, func(i, j int) bool {
		return counts[i][1] < counts[j][1] || (counts[i][1] == counts[j][1] && counts[i][0] < counts[j][0])
	})

	result := []int{}
	for i := 0; i < k; i++ {
		result = append(result, counts[i][0])
	}

	return result
}
