// 231. Power of Two
// https://leetcode.com/problems/power-of-two/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(isPowerOfTwo(1))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(isPowerOfTwo(16))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(isPowerOfTwo(218))
	pp.Println(false)
	pp.Println("=========================================")
	pp.Println(isPowerOfTwo(0))
	pp.Println(false)
	pp.Println("=========================================")
}

func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}

// func isPowerOfTwo(n int) bool {
// 	for ; n > 1 && n&1 != 1; n >>= 1 {
// 	}
// 	return n == 1
// }

// func isPowerOfTwo(n int) bool {
// 	for n > 1 {
// 		if n&1 == 1 {
// 			return false
// 		}
// 		n >>= 1
// 	}
// 	return n == 1
// }
