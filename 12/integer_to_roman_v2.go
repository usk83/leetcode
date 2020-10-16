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

type romanNumeral struct {
	number int
	roman  string
}

var romanNumerals []romanNumeral = []romanNumeral{
	romanNumeral{
		number: 1000,
		roman:  "M",
	},
	romanNumeral{
		number: 900,
		roman:  "CM",
	},
	romanNumeral{
		number: 500,
		roman:  "D",
	},
	romanNumeral{
		number: 400,
		roman:  "CD",
	},
	romanNumeral{
		number: 100,
		roman:  "C",
	},
	romanNumeral{
		number: 90,
		roman:  "XC",
	},
	romanNumeral{
		number: 50,
		roman:  "L",
	},
	romanNumeral{
		number: 40,
		roman:  "XL",
	},
	romanNumeral{
		number: 10,
		roman:  "X",
	},
	romanNumeral{
		number: 9,
		roman:  "IX",
	},
	romanNumeral{
		number: 5,
		roman:  "V",
	},
	romanNumeral{
		number: 4,
		roman:  "IV",
	},
	romanNumeral{
		number: 1,
		roman:  "I",
	},
}

func intToRoman(num int) string {
	roman := ""
	for _, romanNumeral := range romanNumerals {
		for num >= romanNumeral.number {
			roman += romanNumeral.roman
			num -= romanNumeral.number
		}
	}
	return roman
}
