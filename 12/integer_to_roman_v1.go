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
}

// The first stupid solution
func intToRoman(num int) string {
	roman := ""
	for num >= 1000 {
		roman += "M"
		num -= 1000
	}

	if num >= 900 {
		roman += "CM"
		num -= 900
	}
	if num >= 500 {
		roman += "D"
		num -= 500
		for num >= 100 {
			roman += "C"
			num -= 100
		}
	}
	if num >= 400 {
		roman += "CD"
		num -= 400
	}
	for num >= 100 {
		roman += "C"
		num -= 100
	}

	if num >= 90 {
		roman += "XC"
		num -= 90
	}
	if num >= 50 {
		roman += "L"
		num -= 50
		for num >= 10 {
			roman += "X"
			num -= 10
		}
	}
	if num >= 40 {
		roman += "XL"
		num -= 40
	}
	for num >= 10 {
		roman += "X"
		num -= 10
	}

	if num >= 9 {
		roman += "IX"
		num -= 9
	}
	if num >= 5 {
		roman += "V"
		num -= 5
		for num >= 1 {
			roman += "I"
			num -= 1
		}
	}
	if num >= 4 {
		roman += "IV"
		num -= 4
	}
	for num >= 1 {
		roman += "I"
		num -= 1
	}

	return roman
}
