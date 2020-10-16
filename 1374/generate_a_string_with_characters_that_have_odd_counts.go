// 1374. Generate a String With Characters That Have Odd Counts
// https://leetcode.com/problems/generate-a-string-with-characters-that-have-odd-counts/
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
	bytes := make([]byte, n)
	for i := range bytes[1:] {
		bytes[i] = 'a'
	}
	if n&1 == 1 {
		bytes[n-1] = 'a'
	} else {
		bytes[n-1] = 'b'
	}
	return string(bytes)
}
