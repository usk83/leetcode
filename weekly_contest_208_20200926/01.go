// 5523. Crawler Log Folder
// https://leetcode.com/contest/weekly-contest-208/problems/crawler-log-folder/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(minOperations([]string{"d1/", "d2/", "../", "d21/", "./"}))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(minOperations([]string{"d1/", "d2/", "./", "d3/", "../", "d31/"}))
	pp.Println(3)
	pp.Println("=========================================")
	pp.Println(minOperations([]string{"d1/", "../", "../", "../"}))
	pp.Println(0)
	pp.Println("=========================================")
}

func minOperations(logs []string) int {
	var count int
	for _, log := range logs {
		switch log {
		case "./":
		case "../":
			count--
		default:
			count++
		}
		if count < 0 {
			count = 0
		}
	}
	return count
}
