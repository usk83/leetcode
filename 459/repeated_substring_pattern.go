// 459. Repeated Substring Pattern
// https://leetcode.com/problems/repeated-substring-pattern/
package main

import (
	"strings"

	"github.com/k0kubun/pp"
)

// September LeetCoding Challenge Day 3
// https://leetcode.com/explore/challenge/card/september-leetcoding-challenge/554/week-1-september-1st-september-7th/3447/
func main() {
	pp.Println("=========================================")
	pp.Println(repeatedSubstringPattern("abab"))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(repeatedSubstringPattern("aba"))
	pp.Println(false)
	pp.Println("=========================================")
	pp.Println(repeatedSubstringPattern("abcabcabcabc"))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(repeatedSubstringPattern("bb"))
	pp.Println(true)
	pp.Println("=========================================")
}

func repeatedSubstringPattern(s string) bool {
	if len(s) == 0 {
		return false
	}
	ss := (s + s)[1 : len(s)*2-1]
	return strings.Contains(ss, s)
}

// func repeatedSubstringPattern(s string) bool {
// LOOP:
// 	for l := len(s) / 2; l > 0; l-- {
// 		if len(s)%l != 0 {
// 			continue
// 		}
// 		for i := len(s)/l - 1; i > 0; i-- {
// 			for j := 0; j < l; j++ {
// 				if s[j] != s[l*i+j] {
// 					continue LOOP
// 				}
// 			}
// 		}
// 		return true
// 	}
// 	return false
// }
