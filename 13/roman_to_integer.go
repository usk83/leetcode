// 13. Roman to Integer
// https://leetcode.com/problems/roman-to-integer/
package main

import "github.com/k0kubun/pp"

func main() {
	pp.Println("=========================================")
	pp.Println(romanToInt("III"))
	pp.Println(3)
	pp.Println("=========================================")
	pp.Println(romanToInt("IV"))
	pp.Println(4)
	pp.Println("=========================================")
	pp.Println(romanToInt("IX"))
	pp.Println(9)
	pp.Println("=========================================")
	pp.Println(romanToInt("LVIII"))
	pp.Println(58)
	pp.Println("=========================================")
	pp.Println(romanToInt("MCMXCIV"))
	pp.Println(1994)
	pp.Println("=========================================")
}

var romanNumerals map[byte]int = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	var sum, prev int
	for i := len(s) - 1; i >= 0; i-- {
		cur := romanNumerals[s[i]]
		if cur < prev {
			sum -= cur
		} else {
			sum += cur
		}
		prev = cur
	}
	return sum
}

// var romanNumerals map[byte]int = map[byte]int{
// 	'I': 1,
// 	'V': 5,
// 	'X': 10,
// 	'L': 50,
// 	'C': 100,
// 	'D': 500,
// 	'M': 1000,
// }

// func romanToInt(s string) int {
// 	if s == "" {
// 		return 0
// 	}
// 	if len(s) == 1 {
// 		return romanNumerals[s[0]]
// 	}
// 	sum := romanNumerals[s[len(s)-1]]
// 	for i := len(s) - 2; i >= 0; i-- {
// 		if romanNumerals[s[i]] < romanNumerals[s[i+1]] {
// 			sum -= romanNumerals[s[i]]
// 		} else {
// 			sum += romanNumerals[s[i]]
// 		}
// 	}
// 	return sum
// }

// var single map[byte]int = map[byte]int{
// 	'I': 1,
// 	'V': 5,
// 	'X': 10,
// 	'L': 50,
// 	'C': 100,
// 	'D': 500,
// 	'M': 1000,
// }

// var double map[string]int = map[string]int{
// 	"IV": 4,
// 	"IX": 9,
// 	"XL": 40,
// 	"XC": 90,
// 	"CD": 400,
// 	"CM": 900,
// }

// func romanToInt(s string) int {
// 	var num int

// 	for s != "" {
// 		if len(s) >= 2 {
// 			if val, found := double[s[0:2]]; found {
// 				s = s[2:]
// 				num += val
// 				continue
// 			}
// 		}
// 		if val, found := single[s[0]]; found {
// 			num += val
// 		}
// 		s = s[1:]
// 	}

// 	return num
// }
