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

func buildArray(target []int, n int) []string {
	ops := []string{}
	for i := 1; i <= n; i++ {
		if len(target) == 0 {
			break
		}
		if target[0] == i {
			ops = append(ops, "Push")
			target = target[1:]
		} else {
			ops = append(ops, "Push")
			ops = append(ops, "Pop")
		}
	}
	return ops
}
