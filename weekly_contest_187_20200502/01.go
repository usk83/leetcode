// 5400. Destination City
// https://leetcode.com/contest/weekly-contest-187/problems/destination-city/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
}

func destCity(paths [][]string) string {
	citys := map[string][]string{}
	for _, path := range paths {
		citys[path[0]] = append(citys[path[0]], path[1])
		if _, ok := citys[path[1]]; !ok {
			citys[path[1]] = []string{}
		}
	}
	for name, city := range citys {
		if len(city) == 0 {
			return name
		}
	}
	return ""
}
