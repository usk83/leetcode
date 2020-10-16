// 1221. Split a String in Balanced Strings
// https://leetcode.com/problems/split-a-string-in-balanced-strings/
package main

import "github.com/k0kubun/pp"

func main() {
	pp.Println(balancedStringSplit("RLRRLLRLRL"))
	pp.Println(4)
	pp.Println(balancedStringSplit("RLLLLRRRLR"))
	pp.Println(3)
	pp.Println(balancedStringSplit("LLLLRRRR"))
	pp.Println(1)
}

// 1 <= s.length <= 1000
// s[i] = 'L' or 'R'
func balancedStringSplit(s string) int {
	var count, balance int
	for _, c := range s {
		switch c {
		case 'R':
			balance++
		case 'L':
			balance--
		}
		if balance == 0 {
			count++
		}
	}
	return count
}
