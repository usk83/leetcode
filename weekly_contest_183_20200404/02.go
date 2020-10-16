// 5377. Number of Steps to Reduce a Number in Binary Representation to One
// https://leetcode.com/problems/xxx/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	// pp.Println("=========================================")
	// pp.Println(numSteps("1101"))
	// pp.Println(6)
	// pp.Println("=========================================")
	// pp.Println(numSteps("10"))
	// pp.Println(1)
	// pp.Println("=========================================")
	// pp.Println(numSteps("1"))
	// pp.Println(0)
	// pp.Println("=========================================")
	// pp.Println(numSteps("1111110011101010110011100100101110010100101110111010111110110010"))
	// pp.Println()
	pp.Println("=========================================")
	pp.Println(numSteps("10111100001010011010110110100010111111100100101100111101111000100011110110011111001011011111110000001111101010101101101111001010000111010001100111100011110101001001110110111011001000011000011000000100010110011110101100100011010010100001011100010001100110011010010111110111010010111001111110111110011011100111010100010100101011000110010001010011100111001011101101100100000000001111111101101011101111000101110011111111110111011100011000110010100000001001011100000110000001110010110101111110110000011010"))
	pp.Println()
	pp.Println("=========================================")
}

func numSteps(s string) int {
	var count int
	for s[len(s)-1] == '0' {
		count++
		s = s[:len(s)-1]
	}
	if s == "1" {
		return count
	}
	var zero int
	for i := range s {
		if s[i] == '0' {
			zero++
		}
	}
	return count + zero + 1 + len(s)
}

// 	// n, _ := strconv.ParseInt(s, 2, 64)
// 	// num := int(n)
// 	// var count int
// 	// for num != 1 {
// 	// 	count++
// 	// 	if num&1 == 1 {
// 	// 		num += 1
// 	// 	} else {
// 	// 		num >>= 1
// 	// 	}
// 	// }
// 	// return count
// }
