// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	fmt.Println(groupThePeople([]int{3, 3, 3, 3, 3, 1, 3}))
	fmt.Println([][]int{{5}, {0, 1, 2}, {3, 4, 6}})
	pp.Println("=========================================")
	fmt.Println(groupThePeople([]int{2, 1, 3, 3, 3, 2}))
	fmt.Println([][]int{{1}, {0, 5}, {2, 3, 4}})
	pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
}

// groupSizes.length == n
// 1 <= n <= 500
// 1 <= groupSizes[i] <= n
func groupThePeople(groupSizes []int) [][]int {
	m := map[int][][]int{}

	// 同じ数字のやつは同じところに
	for id, size := range groupSizes {
		if groups, found := m[size]; found {
			for i := range groups {
				if len(groups[i]) == size {
					if i+1 == len(groups) {
						m[size] = append(m[size], []int{id})
					}
					continue
				}
				groups[i] = append(groups[i], id)
				break
			}
		} else {
			m[size] = [][]int{[]int{id}}
		}
	}

	result := [][]int{}
	for i := range m {
		for j := range m[i] {
			result = append(result, m[i][j])
		}
	}

	return result
}
