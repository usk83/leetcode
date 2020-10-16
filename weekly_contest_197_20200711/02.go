// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(numSub("0110111"))
	pp.Println(9)
	pp.Println("=========================================")
	pp.Println(numSub("101"))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(numSub("111111"))
	pp.Println(21)
	pp.Println("=========================================")
	pp.Println(numSub("000"))
	pp.Println(0)
	pp.Println("=========================================")
}

const mod = 1000000007

func numSub(s string) int {
	var count int
	var left int
	for {
		for left < len(s) && s[left] == '0' {
			left++
		}
		if left == len(s) {
			break
		}
		right := left
		for right < len(s)-1 && s[right+1] == '1' {
			right++
		}

		length := right - left + 1
		count = (((length + 1) * length / 2 % mod) + count) % mod
		// pp.Println("=== DEBUG START ======================================")
		// pp.Println("left:", left, "right:", right)
		// pp.Println("count:", count)
		// pp.Println("=== DEBUG END ========================================")
		left = right + 1
	}
	return count
}

// func numSub(s string) int {
// 	start := 0
// 	for start < len(s) && end < len(s) {
// 		for start < len(s) {
// 			if s[start] == '0' {
// 				start++
// 			} else {
// 				break
// 			}
// 		}
// 		if start == len(s) {
// 			break
// 		}
// 		start--
// 		end = start + 1
// 		for end < len(s) {
// 			if s[end] == '1' {
// 				end++
// 			} else {
// 				break
// 			}
// 		}
// 		end--
// 		if start < len(s) && end < len(s) {

// 		}
// 	}
// }
