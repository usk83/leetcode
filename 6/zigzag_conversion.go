// 6. ZigZag Conversion
// https://leetcode.com/problems/zigzag-conversion/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(convert("PAYPALISHIRING", 3))
	pp.Println("PAHNAPLSIIGYIR")
	pp.Println("=========================================")
	pp.Println(convert("PAYPALISHIRING", 4))
	pp.Println("PINALSIGYAHRPI")
	pp.Println("=========================================")
	pp.Println(convert("AB", 1))
	pp.Println("AB")
	pp.Println("=========================================")
	pp.Println(convert("ABC", 2))
	pp.Println("ACB")
	pp.Println("=========================================")
	pp.Println(convert("A", 2))
	pp.Println("A")
	pp.Println("=========================================")
}

// PAYPALISHIRING

// P   A   H   N
// A P L S I I G
// Y   I   R

// P   A   H   N
//  A P L S I I G
//   Y   I   R

// PAHN
// APLSIIG
// YIR

// a  g  m
// b fh ln
// ce ik o
// d  j  p

// 1 1
// 2 2
// 3 4
// 4 6

// func convert(s string, numRows int) string {
// 	if numRows == 1 || len(s) < numRows {
// 		return s
// 	}
// 	chars := make([]byte, 0, len(s))
// LOOP:
// 	for i := 0; i < numRows; i++ {
// 		first, second := (numRows-i-1)<<1, i*2
// 		charIndex := i
// 		for {
// 			steps := []int{}
// 			switch i {
// 			case 0:
// 				steps = []int{first}
// 			case numRows - 1:
// 				steps = []int{second}
// 			default:
// 				steps = []int{first, second}
// 			}

// 			for _, step := range steps {
// 				chars = append(chars, s[charIndex])
// 				charIndex += step
// 				if charIndex >= len(s) {
// 					continue LOOP
// 				}
// 			}
// 		}
// 	}
// 	return string(chars)
// }

func convert(s string, numRows int) string {
	if numRows == 1 || len(s) < numRows {
		return s
	}

	rows := make([][]rune, numRows)
	cur, up := 0, false
	for _, c := range s {
		rows[cur] = append(rows[cur], c)
		if up {
			cur--
			if cur < 0 {
				cur += 2
				up = false
			}
		} else {
			cur++
			if cur == numRows {
				cur -= 2
				up = true
			}
		}
	}

	result := ""
	for _, row := range rows {
		result += string(row)
	}

	return result
}
