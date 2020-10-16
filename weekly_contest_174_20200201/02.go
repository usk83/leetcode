// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
	"sort"

	"github.com/k0kubun/pp"
)

func main() {
	// pp.Println("=========================================")
	// pp.Println(minSetSize([]int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7}))
	// pp.Println(2)
	pp.Println("=========================================")
	pp.Println(minSetSize([]int{7, 7, 7, 7, 7, 7}))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(minSetSize([]int{1, 9}))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(minSetSize([]int{1000, 1000, 3, 7}))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(minSetSize([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	pp.Println(5)
	pp.Println("=========================================")
}

func minSetSize(arr []int) int {
	m := map[int]int{}
	for _, a := range arr {
		m[a]++
	}

	pairs := []int{}
	for _, v := range m {
		pairs = append(pairs, v)
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i] > pairs[j]

	})

	count := 0
	sum := 0
	half := (len(arr) + 1) / 2
	for sum < half {
		count++
		sum += pairs[0]
		pairs = pairs[1:]
	}

	return count
}
