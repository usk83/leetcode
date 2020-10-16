// 5352. Generate a String With Characters That Have Odd Counts
// https://leetcode.com/contest/weekly-contest-179/problems/generate-a-string-with-characters-that-have-odd-counts/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println("No local tests available.")
	pp.Println("=========================================")
}

func generateTheString(n int) string {
	ans := ""
	if n&1 == 1 {
		for i := 0; i < n; i++ {
			ans += "a"
		}
		return ans
	}
	for i := 0; i < n-1; i++ {
		ans += "a"
	}
	ans += "b"
	return ans
}
