// 5368. Find Lucky Integer in an Array
// https://leetcode.com/contest/weekly-contest-182/problems/find-lucky-integer-in-an-array/
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

func findLucky(arr []int) int {
	m := map[int]int{}
	for _, num := range arr {
		m[num]++
	}

	max := -1
	for num, count := range m {
		if num == count {
			if num > max {
				max = num
			}
		}
	}
	return max
}
