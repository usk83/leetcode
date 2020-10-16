// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
	"sort"
	"strconv"

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

func displayTable(data [][]string) [][]string {
	foodsMap := map[string]bool{}
	orders := map[int]map[string]int{}

	for _, order := range data {
		num, _ := strconv.Atoi(order[1])
		foodsMap[order[2]] = true
		if _, ok := orders[num]; !ok {
			orders[num] = map[string]int{}
		}
		orders[num][order[2]]++
	}

	tableNumbers := []int{}
	for num := range orders {
		tableNumbers = append(tableNumbers, num)
	}
	sort.Ints(tableNumbers)

	foods := []string{}
	for f := range foodsMap {
		foods = append(foods, f)
	}
	sort.Strings(foods)

	table := [][]string{}
	table = append(table, []string{"Table"})
	table[0] = append(table[0], foods...)
	for _, tn := range tableNumbers {
		table = append(table, []string{strconv.Itoa(tn)})
		for _, food := range foods {
			table[len(table)-1] = append(table[len(table)-1], strconv.Itoa(orders[tn][food]))
		}
	}

	return table
}
