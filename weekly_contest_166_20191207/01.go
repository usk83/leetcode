// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
	"strconv"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(subtractProductAndSum(234))
	pp.Println(15)
	pp.Println("=========================================")
	pp.Println(subtractProductAndSum(4421))
	pp.Println(21)
	pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
}

func subtractProductAndSum(n int) int {
	numStr := strconv.Itoa(n)

	product, sum := 1, 0
	for index := range numStr {
		num, _ := strconv.Atoi(string(numStr[index]))
		product *= num
		sum += num
	}

	return product - sum
}
