// 12. Integer to Roman
// https://leetcode.com/problems/integer-to-roman/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(intToRoman(3))
	pp.Println("III")
	pp.Println("=========================================")
	pp.Println(intToRoman(4))
	pp.Println("IV")
	pp.Println("=========================================")
	pp.Println(intToRoman(9))
	pp.Println("IX")
	pp.Println("=========================================")
	pp.Println(intToRoman(58))
	pp.Println("LVIII")
	pp.Println("=========================================")
	pp.Println(intToRoman(1994))
	pp.Println("MCMXCIV")
	pp.Println("=========================================")
	pp.Println(intToRoman(3888))
	pp.Println("MMMDCCCLXXXVIII")
	pp.Println("=========================================")
}

var romanNumerals [4][2]byte = [4][2]byte{
	{'I', 'V'},
	{'X', 'L'},
	{'C', 'D'},
	{'M', ' '},
}

func intToRoman(num int) string {
	digits := make([]int, 0, 4)
	for num != 0 {
		digits = append(digits, num%10)
		num /= 10
	}
	romans := make([]byte, 0, 15)
	for i := len(digits) - 1; i >= 0; i-- {
		switch digits[i] {
		case 9:
			romans = append(romans, romanNumerals[i][0], romanNumerals[i+1][0])
		case 4:
			romans = append(romans, romanNumerals[i][0], romanNumerals[i][1])
		default:
			if digits[i] >= 5 {
				romans = append(romans, romanNumerals[i][1])
				digits[i] -= 5
			}
			for digits[i] > 0 {
				romans = append(romans, romanNumerals[i][0])
				digits[i] -= 1
			}
		}
	}
	return string(romans)
}
