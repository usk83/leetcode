// 1009. Complement of Base 10 Integer
// https://leetcode.com/problems/complement-of-base-10-integer/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(bitwiseComplement(5))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(bitwiseComplement(7))
	pp.Println(0)
	pp.Println("=========================================")
	pp.Println(bitwiseComplement(10))
	pp.Println(5)
	pp.Println("=========================================")
	pp.Println(bitwiseComplement(0))
	pp.Println(1)
	pp.Println("=========================================")
}

func bitwiseComplement(N int) int {
	mask := 1
	for mask < N {
		mask = mask<<1 | 1
	}
	return N ^ mask
}
