// 5367. Longest Happy Prefix
// https://leetcode.com/contest/weekly-contest-181/problems/longest-happy-prefix/
package main

import (
	"strings"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(longestPrefix("level"))
	pp.Println("l")
	pp.Println("=========================================")
	pp.Println(longestPrefix("ababab"))
	pp.Println("abab")
	pp.Println("=========================================")
	pp.Println(longestPrefix("leetcodeleet"))
	pp.Println("leet")
	pp.Println("=========================================")
	pp.Println(longestPrefix("a"))
	pp.Println("")
	pp.Println("=========================================")
	pp.Println(longestPrefix("babbb"))
	pp.Println("b")
	pp.Println("=========================================")
}

func longestPrefix(s string) string {
	if len(s) < 2 {
		return ""
	}
	pointer := 1
	var index int
	for index = strings.Index(s[pointer:], s[:1]); index != -1; index = strings.Index(s[pointer:], s[:1]) {
		length := len(s) - index - pointer
		if s[:length] == s[pointer+index:] {
			return s[:length]
		}
		pointer += index + 1
	}
	return ""
}
