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

func isPathCrossing(path string) bool {
	visited := map[[2]int]bool{}
	var hor, ver int
	visited[[2]int{hor, ver}] = true
	for _, dist := range path {
		switch dist {
		case 'N':
			ver--
		case 'S':
			ver++
		case 'E':
			hor++
		case 'W':
			hor--
		}
		if visited[[2]int{hor, ver}] {
			return true
		}
		visited[[2]int{hor, ver}] = true
	}
	return false
}
