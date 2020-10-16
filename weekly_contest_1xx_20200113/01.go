// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
	"strconv"

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
}

func maximum69Number(num int) int {
	str := strconv.Itoa(num)
	changed := false
	newStr := ""
	for _, c := range str {
		if changed {
			newStr += string(c)
		} else {
			if c == '6' {
				newStr += "9"
				changed = true
			} else {
				newStr += string(c)
			}
		}
	}
	big, _ := strconv.Atoi(newStr)
	return big
}
