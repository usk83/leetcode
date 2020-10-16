// 191. Number of 1 Bits
// https://leetcode.com/problems/number-of-1-bits/
package main

import (
	"strconv"

	"github.com/k0kubun/pp"
)

func main() {
	var num uint64
	pp.Println("=========================================")
	num, _ = strconv.ParseUint("00000000000000000000000000001011", 2, 32)
	pp.Println(hammingWeight(uint32(num)))
	pp.Println(3)
	pp.Println("=========================================")
	num, _ = strconv.ParseUint("00000000000000000000000010000000", 2, 32)
	pp.Println(hammingWeight(uint32(num)))
	pp.Println(1)
	pp.Println("=========================================")
	num, _ = strconv.ParseUint("11111111111111111111111111111101", 2, 32)
	pp.Println(hammingWeight(uint32(num)))
	pp.Println(31)
	pp.Println("=========================================")
}

func hammingWeight(num uint32) int {
	var count uint32
	for num != 0 {
		count += num & 1
		num >>= 1
	}
	return int(count)
}

// func hammingWeight(num uint32) int {
// 	return bits.OnesCount32(num)
// }
